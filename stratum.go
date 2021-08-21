package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethash/eminer/ethash"
	"github.com/ethash/eminer/stratum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

// Stratum mode
func Stratum(stopChan <-chan struct{}) {
	defer func() {
		if r := recover(); r != nil {
			stack := stack(3)
			log.Error(fmt.Sprintf("Recovered in stratum mode -> %s\n%s\n", r, stack))
			go Stratum(stopChan)
		}
	}()

	sc, err := stratum.New(*flagstratum, 4*time.Minute, version, *flagusername, *flagpassword, *flagworkername, false)
	if err != nil {
		log.Crit("Stratum server critical error", "error", err.Error())
	}

	err = sc.Dial()
	if err != nil {
		log.Crit("Stratum server critical error", "error", err.Error())
	}

	w, err := getWork(sc)
	if err != nil {
		log.Crit(err.Error())
	}

	deviceIds := []int{}
	if *flagmine == "all" {
		deviceIds = getAllDevices()
	} else {
		deviceIds = argToIntSlice(*flagmine)
	}

	miner := ethash.NewCL(deviceIds, *flagworkername, version)

	miner.Lock()
	miner.Work = w
	miner.Unlock()

	miner.SetDAGIntensity(*flagdagintensity)

	if *flagkernel != "" {
		miner.SetKernel(argToIntSlice(*flagkernel))
	}

	if *flagintensity != "" {
		miner.SetIntensity(argToIntSlice(*flagintensity))
	}

	err = miner.InitCL()
	if err != nil {
		log.Crit("OpenCL init not completed", "error", err.Error())
	}

	if *flagfan != "" {
		miner.SetFanPercent(argToIntSlice(*flagfan))
	}

	if *flagengineclock != "" {
		miner.SetEngineClock(argToIntSlice(*flagengineclock))
	}

	if *flagmemoryclock != "" {
		miner.SetMemoryClock(argToIntSlice(*flagmemoryclock))
	}

	changeDAG := make(chan struct{})

	stopCheckNewWork := make(chan struct{})
	checkNewWork := func() {
		workCh := make(chan *json.RawMessage)

		method := "notify_Work"

		sc.RegisterNotification(method, workCh)

		var result *json.RawMessage

		for {
			select {
			case <-stopCheckNewWork:
				return
			case result = <-workCh:
				wt, errc := notifyWork(result)
				if errc != nil {
					log.Error("New work error", "error", errc.Error())
					continue
				}

				if miner.Work.ExtraNonce != wt.ExtraNonce {
					log.Debug("Extra nonce changed", "nonce", wt.ExtraNonce)
					miner.Lock()
					miner.Work.ExtraNonce = wt.ExtraNonce
					miner.Work.SizeBits = wt.SizeBits
					miner.Unlock()
				}

				if wt.Target256.Cmp(miner.Work.Target256) != 0 {
					log.Debug("Difficulty changed", "current", fmt.Sprintf("%.3f GH", float64(miner.Work.Difficulty().Uint64())/1e9),
						"new", fmt.Sprintf("%.3f GH", float64(wt.Difficulty().Uint64())/1e9))
					miner.Lock()
					miner.Work.Target256 = wt.Target256
					miner.Unlock()
				}

				if !bytes.Equal(wt.HeaderHash.Bytes(), miner.Work.HeaderHash.Bytes()) {
					log.Info("Work changed, new work", "hash", wt.HeaderHash.TerminalString(), "difficulty", fmt.Sprintf("%.3f GH", float64(wt.Difficulty().Uint64())/1e9))
					if miner.CmpDagSize(wt) {
						log.Warn("DAG size changed, new DAG will be generate")
						changeDAG <- struct{}{}
					}
					miner.Lock()
					miner.Work = wt
					miner.Unlock()
				}
			}
		}
	}

	stopShareInfo := make(chan struct{})
	shareInfo := func() {
		loops := 0
		for {
			select {
			case <-stopShareInfo:
				return
			case <-time.After(30 * time.Second):
				miner.Poll()
				for _, deviceID := range deviceIds {
					log.Info("GPU device information", "device", deviceID,
						"hashrate", fmt.Sprintf("%.3f Mh/s", miner.GetHashrate(deviceID)/1e6),
						"temperature", fmt.Sprintf("%.2f C", miner.GetTemperature(deviceID)),
						"fan", fmt.Sprintf("%.2f%%", miner.GetFanPercent(deviceID)))
				}

				loops++
				if (loops % 6) == 0 {
					log.Info("Mining global report", "solutions", miner.FoundSolutions.Count(), "rejected", miner.RejectedSolutions.Count(),
						"hashrate", fmt.Sprintf("%.3f Mh/s", miner.TotalHashRate()/1e6))
				}
			}
		}
	}

	stopReportHashRate := make(chan struct{})
	reportHashRate := func() {
		randomID := randomHash()
		for {
			select {
			case <-stopReportHashRate:
				return
			case <-time.After(30 * time.Second):
				hashrate := hexutil.Uint64(uint64(miner.TotalHashRate()))
				params := []interface{}{
					hashrate,
					randomID,
				}
				_, err = sc.SubmitHashrate(params)
				if err != nil {
					log.Error("Submit hash rate error", "error", err.Error())
				}
			}
		}
	}

	log.Info("New work from network", "hash", w.HeaderHash.TerminalString(), "difficulty", fmt.Sprintf("%.3f GH", float64(w.Difficulty().Uint64())/1e9))
	log.Info("Starting mining process", "hash", miner.Work.HeaderHash.TerminalString())

	var wg sync.WaitGroup
	stopFarmMine := make(chan struct{}, len(deviceIds))
	for _, deviceID := range deviceIds {
		wg.Add(1)
		go func(deviceID int) {
			defer wg.Done()

			farmMineByDevice(miner, deviceID, sc, stopFarmMine)
		}(deviceID)
	}

	go checkNewWork()
	go shareInfo()
	go reportHashRate()

	if *flaghttp != "" && *flaghttp != "no" {
		httpServer.SetMiner(miner)
	}

	miner.Poll()

	for {
		select {
		case <-stopChan:
			stopCheckNewWork <- struct{}{}
			stopShareInfo <- struct{}{}
			stopReportHashRate <- struct{}{}

			for range deviceIds {
				stopFarmMine <- struct{}{}
			}

			wg.Wait()
			miner.Destroy()

			sc.Close(true)
			return
		case <-changeDAG:
			for range deviceIds {
				stopFarmMine <- struct{}{}
			}

			wg.Wait()

			err = miner.ChangeDAGOnAllDevices()
			if err != nil {
				if strings.Contains(err.Error(), "Device memory may be insufficient") {
					log.Crit("Generate DAG failed", "error", err.Error())
				}

				log.Error("Generate DAG failed", "error", err.Error())

				miner.Destroy()
				httpServer.ClearStats()
				sc.Close(true)

				go Stratum(stopChan)
				return
			}

			miner.Resume()

			for _, deviceID := range deviceIds {
				wg.Add(1)
				go func(deviceID int) {
					defer wg.Done()

					farmMineByDevice(miner, deviceID, sc, stopFarmMine)
				}(deviceID)
			}
		}
	}
}
