package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethash/eminer/ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

// Benchmark mode
func Benchmark(stopChan chan struct{}) {
	deviceID := *flagbenchmark

	miner := ethash.NewCL([]int{int(deviceID)}, *flagworkername, version)

	hh := common.BytesToHash(common.FromHex(randomHash()))
	sh := common.BytesToHash(common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"))
	diff := new(big.Int).SetUint64(9e10)
	work := ethash.NewWork(45, hh, sh, new(big.Int).Div(ethash.MaxUint256, diff), *flagfixediff)

	miner.Work = work

	miner.SetDAGIntensity(*flagdagintensity)

	if *flagkernel != "" {
		miner.SetKernel(argToIntSlice(*flagkernel))
	}

	if *flagintensity != "" {
		miner.SetIntensity(argToIntSlice(*flagintensity))
	}

	err := miner.InitCL()
	if err != nil {
		log.Crit(fmt.Sprintf("OpenCL init error: %s", err.Error()))
		return
	}

	if *flagfan != "" {
		miner.SetFanPercent(argToIntSlice(*flagfan))
	}

	/*if *flagengineclock != "" {
		miner.SetEngineClock(argToIntSlice(*flagengineclock))
	}

	if *flagmemoryclock != "" {
		miner.SetMemoryClock(argToIntSlice(*flagmemoryclock))
	}*/

	stopReportHashRate := make(chan struct{})
	reportHashRate := func() {
		for {
			select {
			case <-stopReportHashRate:
				return
			case <-time.After(5 * time.Second):
				miner.Poll()
				log.Info("GPU device information", "device", deviceID,
					"hashrate", fmt.Sprintf("%.3f Mh/s", miner.GetHashrate(deviceID)/1e6),
					"temperature", fmt.Sprintf("%.2f C", miner.GetTemperature(deviceID)),
					"fan", fmt.Sprintf("%.2f%%", miner.GetFanPercent(deviceID)))
			}
		}
	}

	var wg sync.WaitGroup

	stopSeal := make(chan struct{})
	seal := func() {
		wg.Add(1)
		defer wg.Done()

		onSolutionFound := func(ok bool, nonce uint64, digest []byte, roundVariance uint64) {
			if nonce != 0 && ok {
				log.Debug("Solution details", "digest", common.ToHex(digest), "nonce", hexutil.Uint64(nonce).String())

				miner.FoundSolutions.Update(int64(roundVariance))
				if *flagfixediff {
					formatter := func(x int64) string {
						return fmt.Sprintf("%d%%", x)
					}

					log.Info("Solutions round variance", "count", miner.FoundSolutions.Count(), "last", formatter(int64(roundVariance)),
						"mean", formatter(int64(miner.FoundSolutions.Mean())), "min", formatter(miner.FoundSolutions.Min()),
						"max", formatter(miner.FoundSolutions.Max()))
				}

				miner.Lock()
				miner.Work.HeaderHash = common.BytesToHash(common.FromHex(randomHash()))
				miner.Unlock()
			}
		}

		miner.Seal(stopSeal, deviceID, onSolutionFound)
	}

	log.Info("Starting benchmark", "seconds for", 600)

	go seal()
	go reportHashRate()

	miner.WorkChanged()

	if *flaghttp != "" && *flaghttp != "no" {
		httpServer.SetMiner(miner)
	}

	select {
	case <-stopChan:
		stopReportHashRate <- struct{}{}
		stopSeal <- struct{}{}

		wg.Wait()
		miner.Release(deviceID)

		log.Info("Benchmark aborted", "device", deviceID, "hashrate", fmt.Sprintf("%.3f Mh/s", miner.GetHashrate(deviceID)/1e6))

		return
	case <-time.After(600 * time.Second):
		stopReportHashRate <- struct{}{}
		stopSeal <- struct{}{}

		wg.Wait()
		miner.Release(deviceID)

		log.Info("Benchmark completed", "device", deviceID, "hashrate", fmt.Sprintf("%.3f Mh/s", miner.GetHashrate(deviceID)/1e6))

		stopChan <- struct{}{}
		return
	}
}
