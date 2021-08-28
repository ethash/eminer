// Copyright 2014 The go-ethereum Authors

package ethash

import (
	"bytes"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/ethash/eminer/adl"
	"github.com/ethash/eminer/counter"
	clbin "github.com/ethash/eminer/ethash/cl"
	"github.com/ethash/eminer/ethash/gcn"
	"github.com/ethash/eminer/nvml"
	"github.com/ethash/go-opencl/cl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hako/durafmt"
	metrics "github.com/rcrowley/go-metrics"
)

//OpenCLDevice struct
type OpenCLDevice struct {
	sync.RWMutex

	deviceID int
	device   *cl.Device

	name string

	busNumber int

	openCL11 bool
	openCL12 bool
	openCL20 bool

	amdGPU    bool
	nvidiaGPU bool

	dagBuf1 *cl.MemObject
	dagBuf2 *cl.MemObject

	headerBuf *cl.MemObject // Hash of block-to-mine in device mem

	searchBuffers []*cl.MemObject

	searchKernel *cl.Kernel

	queue        *cl.CommandQueue
	queueWorkers []*cl.CommandQueue

	ctx        *cl.Context
	program    *cl.Program
	dagProgram *cl.Program

	nonceRand *mrand.Rand // seeded by crypto/rand, see comments where it's initialised

	hashRate    metrics.Meter
	temperature metrics.GaugeFloat64
	fanpercent  metrics.GaugeFloat64
	memoryclock metrics.Gauge
	engineclock metrics.Gauge

	kernel         *kernel
	globalWorkSize uint64
	workGroupSize  uint64

	roundCount counter.Counter

	logger log.Logger
}

//OpenCLMiner struct
type OpenCLMiner struct {
	sync.RWMutex

	ethash *Ethash // Ethash full DAG & cache in host mem
	Work   *Work

	deviceIds []int
	devices   []*OpenCLDevice

	kernels   []int
	intensity []int

	dagSize   uint64
	cacheSize uint64

	workerName string

	stop bool

	SolutionsHashRate metrics.Meter
	FoundSolutions    metrics.Histogram
	RejectedSolutions metrics.Counter
	InvalidSolutions  metrics.Counter

	dagIntensity int

	workCh chan struct{}

	version string

	uptime time.Time
}

type search struct {
	bufIndex    uint32
	startNonce  uint64
	headerHash  common.Hash
	workChanged bool
}

type kernel struct {
	id          int
	source      string
	threadCount uint64
}

const (
	sizeOfUint32 = 4
	sizeOfNode   = 64

	maxSearchResults = 7

	maxWorkGroupSize = 256

	amdIntensity    = 16
	nvidiaIntensity = 8

	defaultSearchBufSize = 1
)

var searchBufSize int

var kernels = []*kernel{
	{id: 1, source: kernelSource("kernel1.cl"), threadCount: 4},
	{id: 2, source: kernelSource("kernel2.cl"), threadCount: 8},
	{id: 3, source: kernelSource("kernel3.cl"), threadCount: 2},
}

//NewCL func
func NewCL(deviceIds []int, workerName, version string) *OpenCLMiner {
	ids := make([]int, len(deviceIds))
	copy(ids, deviceIds)

	miner := &OpenCLMiner{
		dagSize:           0,
		deviceIds:         ids,
		workerName:        workerName,
		SolutionsHashRate: metrics.NewMeter(),
		FoundSolutions:    metrics.NewHistogram(metrics.NewUniformSample(1e4)),
		RejectedSolutions: metrics.NewCounter(),
		InvalidSolutions:  metrics.NewCounter(),
		dagIntensity:      8,
		workCh:            make(chan struct{}),
		version:           version,
		uptime:            time.Now(),
	}

	metrics.Register(workerName+".solutions.hashrate", miner.SolutionsHashRate)
	metrics.Register(workerName+".solutions.found", miner.FoundSolutions)
	metrics.Register(workerName+".solutions.rejected", miner.RejectedSolutions)
	metrics.Register(workerName+".solutions.invalid", miner.InvalidSolutions)

	return miner
}

// InitCL func
func (c *OpenCLMiner) InitCL() error {
	platforms, err := cl.GetPlatforms()
	if err != nil {
		return fmt.Errorf("plaform error: %v\ncheck your OpenCL installation and drivers and then run eminer -L", err)
	}

	var devices []*cl.Device
	for _, p := range platforms {
		ds, err := cl.GetDevices(p, cl.DeviceTypeGPU)
		if err != nil {
			continue
		}

		for _, d := range ds {
			if !strings.Contains(d.Vendor(), "AMD") &&
				!strings.Contains(d.Vendor(), "Advanced Micro Devices") &&
				!strings.Contains(d.Vendor(), "NVIDIA") {
				continue
			}
			devices = append(devices, d)
		}
	}

	blockNum := c.Work.BlockNumberU64()

	pow := New("", 1, 0, "", 1, 0)
	//pow.dataset(blockNum) // generates DAG on CPU if we don't have it
	pow.cache(blockNum) // and cache

	c.ethash = pow
	c.dagSize = datasetSize(blockNum)
	c.cacheSize = cacheSize(blockNum)

	searchBufSize = defaultSearchBufSize

	if runtime.GOOS == "windows" {
		searchBufSize = 2
	}

	var wg sync.WaitGroup
	var errd error

	for i, id := range c.deviceIds {
		if id > len(devices)-1 {
			return fmt.Errorf("device id not found. see available device ids with: eminer -L")
		}

		wg.Add(1)
		go func(idx, deviceID int, device *cl.Device) {
			defer wg.Done()
			errd = c.initCLDevice(idx, deviceID, device)
		}(i, id, devices[id])
	}

	wg.Wait()

	if errd != nil {
		return errd
	}

	if len(c.devices) == 0 {
		return fmt.Errorf("no devices found")
	}

	return nil
}

func (c *OpenCLMiner) initCLDevice(idx, deviceID int, device *cl.Device) error {
	logger := log.New("device", deviceID)

	devGlobalMem := uint64(device.GlobalMemSize())

	if device.Version() == "OpenCL 1.0" {
		return fmt.Errorf("opencl version not supported %s", device.Version())
	}

	var cl11, cl12, cl20 bool
	if strings.Contains(device.Version(), "OpenCL 1.1") {
		cl11 = true
	}
	if strings.Contains(device.Version(), "OpenCL 1.2") {
		cl12 = true
	}
	if strings.Contains(device.Version(), "OpenCL 2.0") {
		cl20 = true
	}

	var amdGPU, nvidiaGPU bool
	var vendor string

	if strings.Contains(device.Vendor(), "AMD") || strings.Contains(device.Vendor(), "Advanced Micro Devices") {
		amdGPU = true
		vendor = "AMD"
	} else if strings.Contains(device.Vendor(), "NVIDIA") {
		nvidiaGPU = true
		vendor = "NVIDIA"
	}

	// log warnings but carry on; some device drivers report inaccurate values
	if c.dagSize > devGlobalMem {
		return fmt.Errorf("device memory may be insufficient, max device memory size: %v DAG size: %v", devGlobalMem, c.dagSize)
	}

	context, err := cl.CreateContext([]*cl.Device{device})
	if err != nil {
		return fmt.Errorf("failed creating context: %v", err)
	}

	queue, err := context.CreateCommandQueue(device, 0)
	if err != nil {
		return fmt.Errorf("command queue err: %v", err)
	}

	busNumber := uint(0)
	if amdGPU {
		busNumber, _ = device.DeviceBusAMD()
	} else if nvidiaGPU {
		busNumber, _ = device.DeviceBusNVIDIA()
	}

	var name string
	if amdGPU {
		name = adl.Name(int(busNumber))
	}

	if len(name) <= 0 {
		name = device.Name()
	}

	kernel := kernels[2] //2T kernel

	if nvidiaGPU {
		kernel = kernels[0] //4T kernel
	}

	if strings.Contains(device.Name(), "1080") {
		kernel = kernels[1] //8T kernel
	}

	if idx < len(c.kernels) {
		if c.kernels[idx] > 0 && c.kernels[idx] < 4 {
			kernel = kernels[c.kernels[idx]-1]
		}
	}

	logger.Info("Initialising", "kernel", kernel.id, "name", name, "vendor", vendor,
		"clock", fmt.Sprintf("%d MHz", device.MaxClockFrequency()), "memory", fmt.Sprintf("%d MB", device.GlobalMemSize()/1024/1024))

	var workGroupSize, globalWorkSize uint64
	var intensity int

	if amdGPU {
		intensity = amdIntensity
	} else if nvidiaGPU {
		intensity = nvidiaIntensity
	}

	if idx < len(c.intensity) {
		intensity = c.intensity[idx]
	}

	if intensity < 8 {
		intensity = 8
	}

	if intensity > 64 {
		intensity = 64
	}

	division := float64(intensity) / 16
	// factor := uint64((32 / float64(intensity)) + 0.5)

	workGroupSize = uint64(intensity * 8)
	// globalWorkSize = uint64(math.Exp2(float64(intensity)/division)*float64(workGroupSize)) * factor
	globalWorkSize = uint64(math.Exp2(float64(intensity)/division) * float64(workGroupSize))

	logger.Trace("Intensity", "intensity", intensity, "global", globalWorkSize, "local", workGroupSize, "bufsize", searchBufSize)

	if workGroupSize > maxWorkGroupSize {
		workGroupSize = maxWorkGroupSize
	}

	globalWorkSize = globalWorkSize / uint64(searchBufSize)

	searchBuffers := make([]*cl.MemObject, searchBufSize)
	for i := 0; i < searchBufSize; i++ {
		searchBuff, errsb := context.CreateEmptyBuffer(cl.MemWriteOnly, (1+maxSearchResults)*sizeOfUint32)
		if errsb != nil {
			return fmt.Errorf("search buffer err: %v", errsb)
		}
		searchBuffers[i] = searchBuff
	}

	queueWorkers := make([]*cl.CommandQueue, searchBufSize)
	for i := 0; i < searchBufSize; i++ {
		queueExec, errq := context.CreateCommandQueue(device, 0)
		if errq != nil {
			return fmt.Errorf("command queue err: %v", errq)
		}
		queueWorkers[i] = queueExec
	}

	headerBuf, err := context.CreateEmptyBuffer(cl.MemReadOnly, 32)
	if err != nil {
		return fmt.Errorf("header buffer err: %v", err)
	}

	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return err
	}
	nonceRand := mrand.New(mrand.NewSource(seed.Int64()))

	d := &OpenCLDevice{
		deviceID: deviceID,
		device:   device,

		name: name,

		busNumber: int(busNumber),

		openCL11: cl11,
		openCL12: cl12,
		openCL20: cl20,

		amdGPU:    amdGPU,
		nvidiaGPU: nvidiaGPU,

		headerBuf:     headerBuf,
		searchBuffers: searchBuffers,

		queue:        queue,
		queueWorkers: queueWorkers,
		ctx:          context,

		nonceRand: nonceRand,

		hashRate:    metrics.NewMeter(),
		temperature: metrics.NewGaugeFloat64(),
		fanpercent:  metrics.NewGaugeFloat64(),
		memoryclock: metrics.NewGauge(),
		engineclock: metrics.NewGauge(),

		kernel:         kernel,
		workGroupSize:  workGroupSize,
		globalWorkSize: globalWorkSize,

		logger: logger,
	}

	metrics.Register(fmt.Sprintf("%s.gpu.%d.hashrate", c.workerName, deviceID), d.hashRate)
	metrics.Register(fmt.Sprintf("%s.gpu.%d.temperature", c.workerName, deviceID), d.temperature)
	metrics.Register(fmt.Sprintf("%s.gpu.%d.fanpercent", c.workerName, deviceID), d.fanpercent)
	metrics.Register(fmt.Sprintf("%s.gpu.%d.memoryclock", c.workerName, deviceID), d.memoryclock)
	metrics.Register(fmt.Sprintf("%s.gpu.%d.engineclock", c.workerName, deviceID), d.engineclock)

	err = c.createDagProgramOnDevice(d)
	if err != nil {
		return err
	}

	err = c.createBinaryProgramOnDevice(d, workGroupSize)
	if err != nil {
		return err
	}

	err = c.generateDAGOnDevice(d)
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()
	c.devices = append(c.devices, d)

	return nil
}

func (c *OpenCLMiner) createBinaryProgramOnDevice(d *OpenCLDevice, workGroupSize uint64) (err error) {
	data, err := gcnSource(fmt.Sprintf("ethash_ellesmere_lws%d.bin", workGroupSize))
	if err != nil {
		return err
	}

	d.program, err = d.ctx.CreateProgramWithBinary(data, d.device)
	if err != nil {
		return fmt.Errorf("program err: %v", err)
	}

	buildOpts := fmt.Sprintf("-D FAST_EXIT=%d", 0)
	err = d.program.BuildProgram([]*cl.Device{d.device}, buildOpts)
	if err != nil {
		return fmt.Errorf("program build err: %v", err)
	}

	d.searchKernel, err = d.program.CreateKernel("search")
	if err != nil {
		return
	}

	return nil
}

func (c *OpenCLMiner) createDagProgramOnDevice(d *OpenCLDevice) (err error) {
	deviceVendor := 0
	if d.nvidiaGPU {
		deviceVendor = 1
	}

	kvs := make(map[string]string)
	kvs["GROUP_SIZE"] = strconv.FormatUint(d.workGroupSize, 10)
	kvs["DEVICE_VENDOR"] = strconv.Itoa(deviceVendor)
	kvs["ACCESSES"] = strconv.FormatUint(loopAccesses, 10)
	kvs["MAX_OUTPUTS"] = strconv.FormatUint(maxSearchResults, 10)
	kvs["HASH_SIZE"] = strconv.FormatUint(d.workGroupSize/d.kernel.threadCount, 10)
	kvs["THREADS"] = strconv.FormatUint(d.kernel.threadCount, 10)
	kernelCode := replaceWords(d.kernel.source, kvs)

	d.dagProgram, err = d.ctx.CreateProgramWithSource([]string{kernelCode})
	if err != nil {
		return fmt.Errorf("program err: %v", err)
	}

	buildOpts := fmt.Sprintf("-D DAG_SIZE=%d -D LIGHT_SIZE=%d", c.dagSize/mixBytes, c.cacheSize/sizeOfNode)
	err = d.program.BuildProgram([]*cl.Device{d.device}, buildOpts)
	if err != nil {
		return fmt.Errorf("program build err: %v", err)
	}

	return nil
}

func (c *OpenCLMiner) generateDAGOnDevice(d *OpenCLDevice) error {
	devGlobalMem := uint64(d.device.GlobalMemSize())
	if c.dagSize > devGlobalMem {
		return fmt.Errorf("device memory may be insufficient, max device memory size: %v DAG size: %v", devGlobalMem, c.dagSize)
	}

	dagKernelFunc := "generate_dag_item"

	dagKernel, err := d.dagProgram.CreateKernel(dagKernelFunc)
	if err != nil {
		return fmt.Errorf("dagKernelName err: %v", err)
	}

	blockNum := c.Work.BlockNumberU64()
	cache := c.ethash.cache(blockNum)

	dagSize1 := c.dagSize / 2
	dagSize2 := c.dagSize / 2

	if c.dagSize/mixBytes&1 > 0 {
		dagSize1 = c.dagSize/2 + 64
		dagSize2 = c.dagSize/2 - 64
	}

	d.dagBuf1, err = d.ctx.CreateEmptyBuffer(cl.MemReadOnly, dagSize1)
	if err != nil {
		return fmt.Errorf("allocating dag buf failed: %v", err)
	}

	d.dagBuf2, err = d.ctx.CreateEmptyBuffer(cl.MemReadOnly, dagSize2)
	if err != nil {
		return fmt.Errorf("allocating dag buf failed: %v", err)
	}

	cacheBuf, err := d.ctx.CreateEmptyBuffer(cl.MemReadOnly, c.cacheSize)
	if err != nil {
		return fmt.Errorf("cache buffer err: %v", err)
	}

	cachePtr := unsafe.Pointer(&cache[0])
	_, err = d.queue.EnqueueWriteBuffer(cacheBuf, true, 0, c.cacheSize, cachePtr, nil)
	if err != nil {
		return fmt.Errorf("writing to cache buf failed: %v", err)
	}

	factor := float64(c.dagIntensity) / 16
	dagWorkGroupSize := uint64(c.dagIntensity * 8)
	dagGlobalWorkSize := uint64(math.Exp2(float64(c.dagIntensity)/factor)) * dagWorkGroupSize / 8

	work := uint64(c.dagSize / sizeOfNode)
	fullRuns := uint64(work / dagGlobalWorkSize)
	restWork := uint64(work % dagGlobalWorkSize)
	if restWork > 0 {
		fullRuns++
	}

	d.queue.Finish()

	err = dagKernel.SetArg(1, cacheBuf)
	if err != nil {
		return fmt.Errorf("set arg failed %v", err)
	}

	err = dagKernel.SetArg(2, d.dagBuf1)
	if err != nil {
		return fmt.Errorf("set arg failed %v", err)
	}

	err = dagKernel.SetArg(3, d.dagBuf2)
	if err != nil {
		return fmt.Errorf("set arg failed %v", err)
	}

	d.logger.Info("Requiring new DAG on device", "epoch", blockNum/epochLength)

	start := time.Now().UnixNano()

	for i := uint64(0); i < fullRuns; i++ {
		err = dagKernel.SetArg(0, uint32(i*dagGlobalWorkSize))
		if err != nil {
			return fmt.Errorf("set arg failed %v", err)
		}

		_, err = d.queue.EnqueueNDRangeKernel(dagKernel,
			[]int{0},
			[]int{int(dagGlobalWorkSize)},
			[]int{int(dagWorkGroupSize)}, nil)
		if err != nil {
			return fmt.Errorf("enqueue dag kernel failed %v", err)
		}

		err = d.queue.Finish()
		if err != nil {
			return fmt.Errorf("clFinish dag queue failed %v", err)
		}

		elapsed := time.Now().UnixNano() - start
		d.logger.Debug("Generating DAG in progress", "epoch", blockNum/epochLength,
			"percentage", int(100*float64(i+1)/float64(fullRuns)), "elapsed", common.PrettyDuration(elapsed))
	}

	elapsed := time.Now().UnixNano() - start
	d.logger.Info("Generated DAG on device", "epoch", blockNum/epochLength, "elapsed", common.PrettyDuration(elapsed))

	cacheBuf.Release()
	dagKernel.Release()

	return nil
}

// ChangeDAGOnAllDevices generate dag on all devices
func (c *OpenCLMiner) ChangeDAGOnAllDevices() (err error) {
	blockNum := c.Work.BlockNumberU64()

	c.dagSize = datasetSize(blockNum)
	c.cacheSize = cacheSize(blockNum)

	var wg sync.WaitGroup

	for _, d := range c.devices {
		d.dagBuf1.Release()
		d.dagBuf2.Release()
		d.searchKernel.Release()
		d.program.Release()

		err = c.createDagProgramOnDevice(d)
		if err != nil {
			return
		}

		err = c.createBinaryProgramOnDevice(d, d.workGroupSize)
		if err != nil {
			return
		}

		wg.Add(1)
		go func(d *OpenCLDevice) {
			defer wg.Done()
			err = c.generateDAGOnDevice(d)
		}(d)
	}

	wg.Wait()

	return
}

// Destroy miner
func (c *OpenCLMiner) Destroy() {
	c.ReleaseAll()
	metrics.DefaultRegistry.UnregisterAll()
}

// ReleaseAll device
func (c *OpenCLMiner) ReleaseAll() {
	log.Info("Releasing all OpenCL devices")

	for _, d := range c.devices {
		d.ctx.Release()
		d.queue.Release()
		d.program.Release()
		d.searchKernel.Release()
		d.dagBuf1.Release()
		d.dagBuf2.Release()
		d.headerBuf.Release()
		for _, q := range d.queueWorkers {
			q.Release()
		}
		for _, s := range d.searchBuffers {
			s.Release()
		}
	}
}

//Release selected device
func (c *OpenCLMiner) Release(deviceID int) {
	index := c.getDevice(deviceID)
	d := c.devices[index]

	d.logger.Info("Releasing device", "name", d.name)

	d.ctx.Release()
	d.queue.Release()
	d.program.Release()
	d.searchKernel.Release()
	d.dagBuf1.Release()
	d.dagBuf2.Release()
	d.headerBuf.Release()
	for _, q := range d.queueWorkers {
		q.Release()
	}
	for _, s := range d.searchBuffers {
		s.Release()
	}
}

// CmpDagSize based on block number
func (c *OpenCLMiner) CmpDagSize(work *Work) bool {
	newDagSize := datasetSize(work.BlockNumberU64())

	return newDagSize != c.dagSize
}

// Seal hashes on GPU
func (c *OpenCLMiner) Seal(stop <-chan struct{}, deviceID int, onSolutionFound func(bool, uint64, []byte, uint64)) error {

	//may stop requested
	time.Sleep(1 * time.Millisecond)
	select {
	case <-stop:
		return nil
	default:
	}

	c.Lock()
	headerHash := c.Work.HeaderHash
	target256 := new(big.Int).SetBytes(c.Work.Target256.Bytes())
	minerTarget := c.Work.MinerTarget
	extraNonce := c.Work.ExtraNonce

	target64 := new(big.Int).Rsh(minerTarget, 192).Uint64()
	if !c.Work.FixedDifficulty {
		target64 = new(big.Int).Rsh(target256, 192).Uint64()
	}
	c.Unlock()

	var zero uint32

	idx := c.getDevice(deviceID)
	d := c.devices[idx]

	var minDeviceRand, maxDeviceRand int64
	segDevice := math.MaxInt64 / int64(len(c.devices))

	if idx == 0 {
		minDeviceRand = 0
	} else {
		minDeviceRand = segDevice * int64(idx)
	}

	maxDeviceRand = segDevice * int64(idx+1)

	_, err := d.queue.EnqueueWriteBuffer(d.headerBuf, false, 0, 32, unsafe.Pointer(&headerHash[0]), nil)
	if err != nil {
		d.logger.Error("Error in seal clEnqueueWriterBuffer", "error", err.Error())
		return err
	}

	for i := 0; i < searchBufSize; i++ {
		_, err = d.queue.EnqueueWriteBuffer(d.searchBuffers[i], false, 0, 4, unsafe.Pointer(&zero), nil)
		if err != nil {
			d.logger.Error("Error in seal clEnqueueWriterBuffer", "error", err.Error())
			return err
		}
	}

	err = d.queue.Finish()
	if err != nil {
		d.logger.Error("Error in seal clFinish", "error", err.Error())
		return err
	}

	err = d.searchKernel.SetArg(1, d.headerBuf)
	if err != nil {
		d.logger.Error("Error in seal clSetKernelArg 1", "error", err.Error())
		return err
	}

	err = d.searchKernel.SetArg(2, d.dagBuf1)
	if err != nil {
		d.logger.Error("Error in seal clSetKernelArg 2", "error", err.Error())
		return err
	}

	err = d.searchKernel.SetArg(3, d.dagBuf2)
	if err != nil {
		d.logger.Error("Error in seal clSetKernelArg 3", "error", err.Error())
		return err
	}

	err = d.searchKernel.SetArg(4, c.dagSize/mixBytes)
	if err != nil {
		d.logger.Error("Error in seal clSetKernelArg 4", "error", err.Error())
		return err
	}

	err = d.searchKernel.SetArg(6, target64)
	if err != nil {
		d.logger.Error("Error in seal clSetKernelArg 6", "error", err.Error())
		return err
	}

	regName := fmt.Sprintf("%s.gpu.%d.hashrate", c.workerName, deviceID)

	metrics.Unregister(regName)
	d.hashRate = metrics.NewMeter()
	metrics.Register(regName, d.hashRate)

	var searchGroup sync.WaitGroup

	worker := func(s *search) {
		runtime.LockOSThread()

		attempts := uint64(0)

		var minWorkerRand, maxWorkerRand int64
		segWorker := (maxDeviceRand - minDeviceRand) / int64(searchBufSize)

		if s.bufIndex == 0 {
			minWorkerRand = minDeviceRand
		} else {
			minWorkerRand = (segWorker * int64(s.bufIndex)) + minDeviceRand
		}

		maxWorkerRand = (segWorker * int64(s.bufIndex+1)) + minDeviceRand

		d.Lock()
		if extraNonce > 0 {
			s.startNonce = extraNonce + (uint64(idx*searchBufSize+int(s.bufIndex)) << (64 - 4 - uint64(c.Work.SizeBits)))
		} else {
			s.startNonce = uint64(d.nonceRand.Int63n(maxWorkerRand-minWorkerRand) + minWorkerRand)
		}
		d.Unlock()

		defer searchGroup.Done()

		var cres *cl.MappedMemObject

		for !c.stop {
			s.headerHash = headerHash

			d.Lock()
			if s.workChanged {
				if extraNonce > 0 {
					s.startNonce = extraNonce + (uint64(idx*searchBufSize+int(s.bufIndex)) << (64 - 4 - uint64(c.Work.SizeBits)))
				} else {
					s.startNonce = uint64(d.nonceRand.Int63n(maxWorkerRand-minWorkerRand) + minWorkerRand)
				}
				s.workChanged = false
			}

			err = d.searchKernel.SetArg(0, d.searchBuffers[s.bufIndex])
			if err != nil {
				d.logger.Error("Error in seal clSetKernelArg 0", "error", err.Error())
				d.Unlock()
				continue
			}

			err = d.searchKernel.SetArg(5, s.startNonce)
			if err != nil {
				d.logger.Error("Error in seal clSetKernelArg 5", "error", err.Error())
				d.Unlock()
				continue
			}

			_, err = d.queueWorkers[s.bufIndex].EnqueueNDRangeKernel(
				d.searchKernel,
				[]int{0},
				[]int{int(d.globalWorkSize)},
				[]int{int(d.workGroupSize)},
				nil)
			if err != nil {
				d.logger.Error("Error in seal clEnqueueNDRangeKernel", "error", err.Error())
				d.Unlock()
				continue
			}
			d.Unlock()

			cres, _, err = d.queueWorkers[s.bufIndex].EnqueueMapBuffer(d.searchBuffers[s.bufIndex], true,
				cl.MapFlagRead, 0, (1+maxSearchResults)*sizeOfUint32,
				nil)
			if err != nil {
				d.logger.Error("Error in seal clEnqueueMapBuffer", "error", err.Error())
				continue
			}

			results := cres.ByteSlice()
			nfound := uint32(math.Min(float64(binary.LittleEndian.Uint32(results)), float64(maxSearchResults)))

			for i := uint32(0); i < nfound; i++ {
				lo := (i + 1) * sizeOfUint32
				hi := (i + 2) * sizeOfUint32
				upperNonce := uint64(binary.LittleEndian.Uint32(results[lo:hi]))
				checkNonce := s.startNonce + upperNonce
				if checkNonce != 0 {
					c.RLock()
					if !bytes.Equal(s.headerHash.Bytes(), c.Work.HeaderHash.Bytes()) {
						d.logger.Warn("Stale solution found", "worker", s.bufIndex,
							"hash", s.headerHash.TerminalString())

						d.roundCount.Empty()

						c.RUnlock()
						continue
					}
					c.RUnlock()

					// We verify that the nonce is indeed a solution by
					// executing the Ethash verification function (on the CPU).
					number := c.Work.BlockNumberU64()
					cache := c.ethash.cache(number)
					mixDigest, result := hashimotoLight(c.dagSize, cache, s.headerHash.Bytes(), checkNonce)

					if new(big.Int).SetBytes(result).Cmp(target256) <= 0 {
						d.logger.Info("Solution found and verified", "worker", s.bufIndex,
							"hash", s.headerHash.TerminalString())

						c.SolutionsHashRate.Mark(c.Work.Difficulty().Int64())

						roundVariance := uint64(100)
						if c.Work.FixedDifficulty {
							d.roundCount.Put()
							roundCount := d.roundCount.Count() * c.Work.MinerDifficulty().Uint64()
							roundVariance = roundCount * 100 / c.Work.Difficulty().Uint64()
						}

						go onSolutionFound(true, checkNonce, mixDigest, roundVariance)

						d.roundCount.Empty()

					} else if c.Work.FixedDifficulty {
						if new(big.Int).SetBytes(result).Cmp(c.Work.MinerTarget) <= 0 {
							d.roundCount.Put()
						}
					} else {
						d.logger.Error("Found corrupt solution, check your device.")
						c.InvalidSolutions.Inc(1)
					}
				}
			}

			if nfound > 0 {
				_, err = d.queueWorkers[s.bufIndex].EnqueueWriteBuffer(d.searchBuffers[s.bufIndex], false, 0, 4, unsafe.Pointer(&zero), nil)
				if err != nil {
					d.logger.Error("Error in seal clEnqueueWriteBuffer", "error", err.Error())
				}
			}

			_, err = d.queueWorkers[s.bufIndex].EnqueueUnmapMemObject(d.searchBuffers[s.bufIndex], cres, nil)
			if err != nil {
				d.logger.Error("Error in seal clEnqueueUnMapMemObject", "error", err.Error())
			}

			d.queueWorkers[s.bufIndex].Finish()

			s.startNonce = s.startNonce + d.globalWorkSize

			attempts++
			if attempts == 2 {
				d.hashRate.Mark(int64(d.globalWorkSize * attempts))
				attempts = 0
			}
		}
	}

	workers := make([]*search, 0, searchBufSize)
	for i := uint32(0); i < uint32(searchBufSize); i++ {
		s := &search{bufIndex: i}
		workers = append(workers, s)
		searchGroup.Add(1)
		go worker(s)
	}

	for {
		select {
		case <-stop:
			c.stop = true
			for _, q := range d.queueWorkers {
				err = q.Finish()
				if err != nil {
					d.logger.Error("Error in seal clFinish", "error", err.Error())
				}
			}

			searchGroup.Wait()

			return err

		case <-c.workCh:
			c.Lock()
			//if the new target > then current one change immediately
			if target256.Cmp(c.Work.Target256) > 0 {
				target256 = new(big.Int).SetBytes(c.Work.Target256.Bytes())

				if !c.Work.FixedDifficulty {
					target64 = new(big.Int).Rsh(target256, 192).Uint64()

					err = d.searchKernel.SetArg(6, target64)
					if err != nil {
						d.logger.Error("Error in seal clSetKernelArg 6", "error", err.Error())
						c.Unlock()
						goto done
					}
				}
			}

			if c.Work.ExtraNonce != extraNonce {
				extraNonce = c.Work.ExtraNonce

				d.Lock()
				for _, s := range workers {
					s.workChanged = true
				}
				d.Unlock()
			}

			if !bytes.Equal(headerHash.Bytes(), c.Work.HeaderHash.Bytes()) {
				headerHash = c.Work.HeaderHash

				_, err = d.queue.EnqueueWriteBuffer(d.headerBuf, false, 0, 32, unsafe.Pointer(&headerHash[0]), nil)
				if err != nil {
					d.logger.Error("Error in seal clEnqueueWriterBuffer", "error", err.Error())
					c.Unlock()
					goto done
				}

				err = d.queue.Finish()
				if err != nil {
					d.logger.Error("Error in seal clFinish", "error", err.Error())
					c.Unlock()
					goto done
				}

				//Some pools doesn't accept solutions with old work like nicehash
				if target256.Cmp(c.Work.Target256) != 0 {
					target256 = new(big.Int).SetBytes(c.Work.Target256.Bytes())

					if !c.Work.FixedDifficulty {
						target64 = new(big.Int).Rsh(target256, 192).Uint64()

						err = d.searchKernel.SetArg(6, target64)
						if err != nil {
							d.logger.Error("Error in seal clSetKernelArg 6", "error", err.Error())
							c.Unlock()
							goto done
						}
					}
				}

				d.Lock()
				for _, s := range workers {
					s.workChanged = true
				}
				d.Unlock()
			}
			c.Unlock()
		}
	}

done:
	c.stop = true

	for _, q := range d.queueWorkers {
		err = q.Finish()
		if err != nil {
			d.logger.Error("Error in seal clFinish", "error", err.Error())
		}
	}

	searchGroup.Wait()

	return err
}

// WorkChanged function
func (c *OpenCLMiner) WorkChanged() {
	c.workCh <- struct{}{}
}

// GetHashrate for device
func (c *OpenCLMiner) GetHashrate(deviceID int) float64 {
	index := c.getDevice(deviceID)
	d := c.devices[index]
	return d.hashRate.RateMean()
}

// GetTemperature for device
func (c *OpenCLMiner) GetTemperature(deviceID int) float64 {
	index := c.getDevice(deviceID)
	d := c.devices[index]
	return d.temperature.Value()
}

// GetFanPercent for device
func (c *OpenCLMiner) GetFanPercent(deviceID int) float64 {
	index := c.getDevice(deviceID)
	d := c.devices[index]
	return d.fanpercent.Value()
}

// TotalHashRate on all GPUs
func (c *OpenCLMiner) TotalHashRate() (total float64) {
	for _, d := range c.devices {
		total += d.hashRate.RateMean()
	}

	return
}

// TotalHashRateMean on all GPUs
func (c *OpenCLMiner) TotalHashRateMean() float64 {
	return c.TotalHashRate()
}

// TotalHashRate1 on all GPUs
func (c *OpenCLMiner) TotalHashRate1() (total float64) {
	for _, d := range c.devices {
		total += d.hashRate.Rate1()
	}

	return
}

// Poll get some useful data from devices
func (c *OpenCLMiner) Poll() {
	for _, d := range c.devices {
		if d.amdGPU {
			d.temperature.Update(adl.Temperature(d.busNumber))
			d.fanpercent.Update(adl.FanPercent(d.busNumber))
			d.engineclock.Update(int64(adl.EngineClock(d.busNumber)))
			d.memoryclock.Update(int64(adl.MemoryClock(d.busNumber)))
		} else if d.nvidiaGPU {
			d.temperature.Update(nvml.Temperature(d.busNumber))
			d.fanpercent.Update(nvml.FanPercent(d.busNumber))
			d.engineclock.Update(int64(nvml.EngineClock(d.busNumber)))
			d.memoryclock.Update(int64(nvml.MemoryClock(d.busNumber)))
		}
	}
}

// SetFanPercent set fan speed percent for selected devices
func (c *OpenCLMiner) SetFanPercent(percents []int) {
	for i, p := range percents {
		if i > len(c.devices)-1 {
			break
		}

		d := c.devices[i]
		if d.amdGPU {
			if err := adl.FanSetPercent(d.busNumber, uint32(p)); err != nil {
				d.logger.Error("Fan set error", "error", err.Error(), "bus", d.busNumber)
			}
		}
	}
}

// SetEngineClock set engine clock for selected devices
func (c *OpenCLMiner) SetEngineClock(values []int) {
	for i, v := range values {
		if i > len(c.devices)-1 {
			break
		}

		d := c.devices[i]
		if d.amdGPU {
			if err := adl.EngineSetClock(d.busNumber, int(v)); err != nil {
				d.logger.Error("Engine clock set error", "error", err.Error(), "bus", d.busNumber)
			}
		}
	}
}

// SetMemoryClock set memory clock for selected devices
func (c *OpenCLMiner) SetMemoryClock(values []int) {
	for i, v := range values {
		if i > len(c.devices)-1 {
			break
		}

		d := c.devices[i]
		if d.amdGPU {
			if err := adl.MemorySetClock(d.busNumber, int(v)); err != nil {
				d.logger.Error("Memory clock set error", "error", err.Error(), "bus", d.busNumber)
			}
		}
	}
}

func (c *OpenCLMiner) getDevice(deviceID int) int {
	for i, d := range c.devices {
		if d.deviceID == deviceID {
			return i
		}
	}
	return 0
}

// SetKernel for each device
func (c *OpenCLMiner) SetKernel(values []int) {
	c.kernels = values
}

// SetDAGIntensity for all devices
func (c *OpenCLMiner) SetDAGIntensity(value int) {
	if value < 4 {
		value = 4
	} else if value > 32 {
		value = 32
	}

	c.dagIntensity = value
}

// SetIntensity for each device
func (c *OpenCLMiner) SetIntensity(values []int) {
	c.intensity = values
}

// LowMemDevice looking low mem devices
func (c *OpenCLMiner) LowMemDevice() bool {
	for _, d := range c.devices {
		if d.device.GlobalMemSize() <= 2*1024*1024*1024 {
			return true
		}
	}

	return false
}

// Resume mining
func (c *OpenCLMiner) Resume() {
	c.stop = false
}

// MarshalJSON for json encoding
func (c *OpenCLMiner) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})

	data["current_work"] = c.Work
	data["worker_name"] = c.workerName
	data["uptime"] = durafmt.Parse(time.Since(c.uptime).Round(time.Second)).String()
	data["uptime_secs"] = time.Since(c.uptime).Seconds()
	data["found_solutions"] = c.FoundSolutions.Count()
	data["rejected_solutions"] = c.RejectedSolutions.Count()
	data["invalid_solutions"] = c.InvalidSolutions.Count()
	data["solutions_hashrate_mean"] = c.SolutionsHashRate.RateMean()
	data["total_hashrate_mean"] = c.TotalHashRateMean()
	data["total_hashrate_1m"] = c.TotalHashRate1()
	data["version"] = c.version

	var devices []map[string]interface{}

	for _, id := range c.deviceIds {
		idx := c.getDevice(id)
		d := c.devices[idx]

		device := make(map[string]interface{})
		device["name"] = d.name
		device["vendor"] = d.device.Vendor()
		device["memory"] = d.device.GlobalMemSize()
		device["max_clock"] = d.device.MaxClockFrequency()
		device["engine_clock"] = d.engineclock.Value()
		device["memory_clock"] = d.memoryclock.Value()
		device["hashrate_mean"] = d.hashRate.RateMean()
		device["hashrate_1m"] = d.hashRate.Rate1()
		device["hashrate_5m"] = d.hashRate.Rate5()
		device["hashrate_15m"] = d.hashRate.Rate15()
		device["temperature"] = d.temperature.Value()
		device["fan_percent"] = d.fanpercent.Value()

		devices = append(devices, device)

	}

	data["devices"] = devices

	return json.Marshal(data)
}

func replaceWords(text string, kvs map[string]string) string {
	for k, v := range kvs {
		text = strings.Replace(text, k, v, -1)
	}
	return text
}

func kernelSource(name string) string {
	asset, err := clbin.Asset("cl/" + name)
	if err != nil {
		return ""
	}

	return string(asset)
}

func gcnSource(name string) ([]byte, error) {
	asset, err := gcn.Asset("gcn/bin/" + name)
	if err != nil {
		return []byte{}, err
	}

	return asset, nil
}
