package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethash/eminer/client"
	"github.com/ethash/eminer/ethash"
	"github.com/ethash/eminer/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func farmMineByDevice(miner *ethash.OpenCLMiner, deviceID int, c client.Client, stopChan <-chan struct{}) {
	stopSealFunc := make(chan struct{})
	stopSeal := make(chan struct{})
	seal := func() {
		for {
			select {
			case <-stopSealFunc:
				return
			case <-time.After(time.Second):
				onSolutionFound := func(ok bool, nonce uint64, digest []byte, roundVariance uint64) {
					blockNonce := types.EncodeNonce(nonce)

					params := []interface{}{
						blockNonce,
						miner.Work.HeaderHash,
						common.BytesToHash(digest),
					}

					miner.FoundSolutions.Update(int64(roundVariance))
					if *flagfixediff {
						formatter := func(x int64) string {
							return fmt.Sprintf("%d%%", x)
						}

						log.Info("Solutions round variance", "count", miner.FoundSolutions.Count(), "last", formatter(int64(roundVariance)),
							"mean", formatter(int64(miner.FoundSolutions.Mean())), "min", formatter(miner.FoundSolutions.Min()),
							"max", formatter(miner.FoundSolutions.Max()))
					}

					res, err := c.SubmitWork(params)
					if res && err == nil {
						log.Info("Solution accepted by network", "hash", miner.Work.HeaderHash.TerminalString())
					} else {
						miner.RejectedSolutions.Inc(1)
						log.Error("Solution not accepted by network", "error", err.Error())
					}
				}

				miner.Seal(stopSeal, deviceID, onSolutionFound)
			}
		}
	}

	go seal()

	<-stopChan

	stopSeal <- struct{}{}
	stopSealFunc <- struct{}{}
}

// Farm mode
func Farm(stopChan <-chan struct{}) {
	defer func() {
		if r := recover(); r != nil {
			stack := stack(3)
			log.Error(fmt.Sprintf("Recovered in farm mode -> %s\n%s\n", r, stack))
			go Farm(stopChan)
		}
	}()

	rc, err := rpc.New(*flagfarm, 2*time.Second)
	if err != nil {
		log.Crit(err.Error())
	}

	w, err := getWork(rc)
	if err != nil {
		if strings.Contains(err.Error(), "No work available") {
			log.Warn("No work available on network, will try 5 sec later")
			time.Sleep(5 * time.Second)
			go Farm(stopChan)
			return
		}
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
		for {
			select {
			case <-stopCheckNewWork:
				return
			case <-time.After(250 * time.Millisecond):
				wt, errc := getWork(rc)
				if errc != nil {
					log.Error("Get work error", "error", errc.Error())
					continue
				}

				if !bytes.Equal(wt.HeaderHash.Bytes(), miner.Work.HeaderHash.Bytes()) {
					log.Info("Work changed, new work", "hash", wt.HeaderHash.TerminalString(), "difficulty",
						fmt.Sprintf("%.3f GH", float64(wt.Difficulty().Uint64())/1e9))

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
				_, err = rc.SubmitHashrate(params)
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

			farmMineByDevice(miner, deviceID, rc, stopFarmMine)
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

				go Farm(stopChan)
				return
			}

			miner.Resume()

			for _, deviceID := range deviceIds {
				wg.Add(1)
				go func(deviceID int) {
					defer wg.Done()

					farmMineByDevice(miner, deviceID, rc, stopFarmMine)
				}(deviceID)
			}
		}
	}
}
