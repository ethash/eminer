package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ethash/eminer/client"
	"github.com/ethash/eminer/ethash"
	"github.com/ethash/go-opencl/cl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func argToIntSlice(arg string) (devices []int) {
	deviceList := strings.Split(arg, ",")

	for _, device := range deviceList {
		deviceID, _ := strconv.Atoi(device)
		devices = append(devices, deviceID)
	}

	return
}

func getAllDevices() (devices []int) {
	platforms, err := cl.GetPlatforms()
	if err != nil {
		log.Crit(fmt.Sprintf("Plaform error: %v\nCheck your OpenCL installation and then run eminer -L", err))
		return
	}

	platformMap := make(map[string]bool, len(platforms))

	found := 0
	for _, p := range platforms {
		// check duplicate platform, sometimes found duplicate platforms
		if platformMap[p.Vendor()] {
			continue
		}

		ds, err := cl.GetDevices(p, cl.DeviceTypeGPU)
		if err != nil {
			continue
		}

		platformMap[p.Vendor()] = true

		for _, d := range ds {
			if !strings.Contains(d.Vendor(), "AMD") &&
				!strings.Contains(d.Vendor(), "Advanced Micro Devices") &&
				!strings.Contains(d.Vendor(), "NVIDIA") {
				continue
			}

			devices = append(devices, found)

			found++
		}
	}

	return
}

func randomHash() string {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 32)
	rand.Read(token)

	return common.ToHex(token)
}

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func notifyWork(result *json.RawMessage) (*ethash.Work, error) {
	var blockNumber int64
	var getWork []string
	err := json.Unmarshal(*result, &getWork)
	if err != nil {
		return nil, err
	}

	if len(getWork) < 3 {
		return nil, errors.New("result short")
	}

	seedHash := common.BytesToHash(common.FromHex(getWork[1]))

	blockNumber, err = ethash.Number(seedHash)
	if err != nil {
		return nil, err
	}

	w := ethash.NewWork(blockNumber, common.BytesToHash(common.FromHex(getWork[0])),
		seedHash, new(big.Int).SetBytes(common.FromHex(getWork[2])), *flagfixediff)

	if len(getWork) > 4 { //extraNonce
		w.ExtraNonce = new(big.Int).SetBytes(common.FromHex(getWork[3])).Uint64()
		w.SizeBits, _ = strconv.Atoi(getWork[4])
	}

	return w, nil
}

func getWork(c client.Client) (*ethash.Work, error) {
	var blockNumber int64

	getWork, err := c.GetWork()
	if err != nil {
		return nil, err
	}

	seedHash := common.BytesToHash(common.FromHex(getWork[1]))

	blockNumber, err = ethash.Number(seedHash)
	if err != nil {
		return nil, err
	}

	w := ethash.NewWork(blockNumber, common.BytesToHash(common.FromHex(getWork[0])),
		seedHash, new(big.Int).SetBytes(common.FromHex(getWork[2])), *flagfixediff)

	if len(getWork) > 4 { //extraNonce
		w.ExtraNonce = new(big.Int).SetBytes(common.FromHex(getWork[3])).Uint64()
		w.SizeBits, _ = strconv.Atoi(getWork[4])
	}

	return w, nil
}
