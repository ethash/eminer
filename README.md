# Eminer ethash miner
Optimized, high performance multiworker ethash miner written in Go language.

### Dashboard
![dashboard](https://raw.githubusercontent.com/ethash/eminer/master/dashboard.png)

## Features
- Fully support AMD and NVIDIA OpenCL devices
- Improved three OpenCL kernels
- Asynchronous multiworker (windows only)
- Support for Stratum and RPC clients with failover
- Useful web dashboard
- Historical metrics for last 24 hours, shares, hashrate and other informations
- JSON API for stats and metrics
- Support for AMD and NVIDIA hardware management (Temperature, fan speed, clock and other useful hw informations)
- Support for nicehash stratum

And much more.

**Multiworker Mode** (windows only); search shares with multiple instances, this can be increase **1% ~ 2%** share luck.

## Running
List Devices:
```
$ eminer -L
```

Benchmark mode:
```
$ eminer -B deviceid
```

Stratum mode:
```
$ eminer -S server:port -U yourwallet -P password 
(for nicehash or other stratum servers use -S stratum+tcp://server:port)
```

HTTP-RPC mode:
```
$ eminer -F http://localhost:8545
```

Usage:
```console
$ eminer -h
Usage of eminer:
  -B int
    	Benchmark mode, set device id for benchmark (default -1)
  -F string
    	Farm mode with the work server at URL, use comma for multiple rpc server 
      (default "http://127.0.0.1:8545")
  -L	List GPU devices
  -M string
    	Run mine on selected devices, use comma for multiple devices (default "all")
  -N string
    	Name of your rig, the name will be use on dashboard, json-api and stathat. 
      Some pools require rig name with extra parameter, this name will be send the pools.
  -P string
    	Password for stratum server
  -S <host>:<port>
    	Stratum mode, use comma for multiple stratum server (example: <host>:<port> 
      for nicehash or other stratum servers stratum+tcp://<host>:<port>)
  -U string
    	Username for stratum server
  -V int
    	Log level (0-5) (default 3)
  -cpu int
    	Set the maximum number of CPUs to use
  -dag-intensity int
    	DAG work size intensity (4-32) (default 32)
  -fan-percent string
    	Set fan speed percent on selected devices, 
      use comma for multiple devices (amd devices only)
  -fixed-diff
    	Fixed diff for works, round solutions
  -http string
    	HTTP server for monitoring (read-only) for disable set "no" (default ":8550")
  -intensity string
    	GPU work size intensity (8-64), use comma for multiple devices (default 32)
  -kernel string
    	Select kernel for GPU devices, currently 3 kernels available, 
      use comma for multiple devices (1-3)
  -no-output-color
    	Disable colorized output log format
  -stathat string
    	Set your stathat email address here to have some basic metrics from stathat.com web site
  -v	Version
```

Run eminer and & view the dashboard at `http://localhost:8550`

### JSON-API Endpoints
- `GET` /api/v1/stats
- `GET` /api/v1/chartData

## TODO
- Alarms and more statistics
- Better support hardware management

## Works 
Dual mining; it has negative profit sometimes, more power consumption and more GPU temperature. Current status, more research and more test.

### Donations ETH
0x4e6f8135f909a943344f065a9ec2bedcc14c750d
