package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ethash/eminer/adl"
	"github.com/ethash/eminer/http"
	"github.com/ethash/eminer/nvml"
	"github.com/ethash/go-opencl/cl"
	"github.com/ethereum/go-ethereum/log"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/stathat"
)

var (
	flaglistdevices  = flag.Bool("L", false, "List GPU devices")
	flagbenchmark    = flag.Int("B", -1, "Benchmark mode, set device id for benchmark")
	flagmine         = flag.String("M", "all", "Run mine on selected devices, use comma for multiple devices")
	flagfarm         = flag.String("F", "http://127.0.0.1:8545", "Farm mode with the work server at URL, use comma for multiple rpc server")
	flagstratum      = flag.String("S", "", "Stratum mode, use comma for multiple stratum server (example: `<host>:<port>` for nicehash or other stratum servers stratum+tcp://<host>:<port>)")
	flagworkername   = flag.String("N", "", "Name of your rig, the name will be use on dashboard, json-api, stathat. Some pools require rig name with extra parameter, this name will be send the pools.")
	flagusername     = flag.String("U", "", "Username for stratum server")
	flagpassword     = flag.String("P", "", "Password for stratum server")
	flagintensity    = flag.String("intensity", "", "GPU work size intensity (8-64), use comma for multiple devices (default 32)")
	flagdagintensity = flag.Int("dag-intensity", 32, "DAG work size intensity (4-32)")
	flagkernel       = flag.String("kernel", "", "Select kernel for GPU devices, currently 3 kernels available, use comma for multiple devices (1-3)")
	flagcpus         = flag.Int("cpu", 0, "Set the maximum number of CPUs to use")
	flaghttp         = flag.String("http", ":8550", "HTTP server for monitoring (read-only) for disable set \"no\"")
	flagfixediff     = flag.Bool("fixed-diff", false, "Fixed diff for works, round solutions")
	flagfan          = flag.String("fan-percent", "", "Set fan speed percent on selected devices, use comma for multiple devices (amd devices only)")
	flagstathat      = flag.String("stathat", "", "Set your stathat email address here to have some basic metrics from stathat.com web site")
	flagnocolor      = flag.Bool("no-output-color", false, "Disable colorized output log format")
	flagloglevel     = flag.Int("V", 3, "Log level (0-5)")
	flagversion      = flag.Bool("v", false, "Version")

	flagengineclock = flag.String("engine-clock", "", "Set engine clock on selected devices, use comma for multiple devices (amd devices only)")
	flagmemoryclock = flag.String("memory-clock", "", "Set memory clock on selected devices, use comma for multiple devices (amd devices only)")

)

var (
	devicesTypesForMining = cl.DeviceTypeGPU
	version               = "0.6.1-rc2"
	httpServer            = &http.Server{}
	virtualTerm           = false
)

// ListDevices list devices from OpenCL
func ListDevices() {

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	show := func(name string, v1 interface{}) {
		fmt.Fprintf(w, "%s\t%v\n", name, v1)
	}

	fmt.Println("")
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("|                  OPENCL DEVICES                  |")
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("")

	var found []*cl.Device

	platforms, err := cl.GetPlatforms()
	if err != nil {
		fmt.Println("Plaform error (check your OpenCL installation):", err)
		return
	}

	fmt.Println("===================  PLATFORMS  ====================")
	fmt.Println("")

	for i, p := range platforms {
		show("Platform ID", i)
		show("Platform Name", p.Name())
		show("Platform Vendor", p.Vendor())
		show("Platform Version", p.Version())
		show("Platform Extensions", p.Extensions())
		show("Platform Profile", p.Profile())
		show("", "")

		devices, err := cl.GetDevices(p, cl.DeviceTypeGPU)
		if err != nil {
			show("No device on current platform:", err)
			continue
		}

		for _, d := range devices {
			if !strings.Contains(d.Vendor(), "AMD") &&
				!strings.Contains(d.Vendor(), "Advanced Micro Devices") &&
				!strings.Contains(d.Vendor(), "NVIDIA") {
				continue
			}

			show("   ===  DEVICE #"+strconv.Itoa(len(found))+"  === ", "")
			show("   Device ID for mining", len(found))
			show("   Device Name ", d.Name())
			show("   Vendor", d.Vendor())
			show("   Version", d.Version())
			show("   Driver version", d.DriverVersion())
			show("   Address bits", d.AddressBits())
			show("   Max clock freq", d.MaxClockFrequency())
			show("   Global mem size", d.GlobalMemSize())
			show("   Max constant buffer size", d.MaxConstantBufferSize())
			show("   Max mem alloc size", d.MaxMemAllocSize())
			show("   Max compute units", d.MaxComputeUnits())
			show("   Max work group size", d.MaxWorkGroupSize())
			show("   Max work item sizes", d.MaxWorkItemSizes())
			if strings.Contains(d.Vendor(), "AMD") ||
				strings.Contains(d.Vendor(), "Advanced Micro Devices") {
				busNumber, _ := d.DeviceBusAMD()
				show("   PCI bus id", busNumber)
			} else if strings.Contains(d.Vendor(), "NVIDIA") {
				busNumber, _ := d.DeviceBusNVIDIA()
				show("   PCI bus id", busNumber)
			}
			show("", "")

			found = append(found, d)
		}

		w.Flush()

		fmt.Println("====================================================")
		fmt.Println("")
	}

	fmt.Println("")

	if len(found) == 0 {
		fmt.Println("No device found. Check that your hardware details.")
	} else {
		var idsFormat string
		for i := 0; i < len(found); i++ {
			idsFormat += strconv.Itoa(i)
			if i != len(found)-1 {
				idsFormat += ","
			}
		}
		fmt.Printf("Found %v devices.", len(found))
	}
	fmt.Println("")
}

func main() {
	flag.Parse()

	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*flagloglevel),
		log.StreamHandler(os.Stdout, log.TerminalFormat(!*flagnocolor))))

	if runtime.GOOS == "windows" {
		log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*flagloglevel),
			log.StreamHandler(os.Stdout, log.TerminalFormat(virtualTerm))))
	}

	if *flagversion {
		fmt.Printf("Version: v%s\n", version)
		return
	}

	if *flagcpus == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	} else {
		runtime.GOMAXPROCS(*flagcpus)
	}

	if *flaglistdevices {
		ListDevices()
		return
	}

	if *flagworkername == "" {
		*flagworkername = randomString(5)
	}

	if *flaghttp != "" && *flaghttp != "no" {
		debug := false
		if *flagloglevel > 3 {
			debug = true
		}

		httpServer = http.New(*flaghttp, version, *flagworkername, debug)
		go httpServer.Start()
	}

	if *flagstathat != "" {
		go stathat.Stathat(metrics.DefaultRegistry, 60e9, *flagstathat)
		log.Info("Stathat registered for metrics", "email", *flagstathat)
	}


	stopChan := make(chan struct{})
	var mode string

	if *flagbenchmark > -1 {
		go Benchmark(stopChan)
		mode = "benchmark"
		goto wait
	}

	if *flagstratum == "" && *flagmine != "" {
		go Farm(stopChan)
		mode = "farm"
		goto wait
	}

	if *flagstratum != "" {
		if *flagusername == "" {
			log.Crit("Stratum username required")
		}

		go Stratum(stopChan)
		mode = "stratum"
		goto wait
	}

	if *flagbenchmark < 0 && *flagmine == "" {
		flag.Usage()
		return
	}

wait:
	log.Info("Starting processes", "worker", *flagworkername, "mode", mode)

	if !nvml.Active() {
		log.Warn("NVIDIA (NVML) hardware monitor library couldn't initialize")
	}

	if !adl.Active() {
		log.Warn("AMD (ADL) hardware monitor library couldn't initialize")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	select {
	case <-signalChan:
		log.Warn("Interrupted by user, processes will be stop")
		stopChan <- struct{}{}
	case <-stopChan:
		return
	}

	time.Sleep(3 * time.Second)

	adl.Release()
	nvml.Release()
}
