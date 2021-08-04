package nvml

/*
#include "nvmlbridge.h"
*/
import "C"

// Device structure
type Device struct {
	nvmldevice C.nvmlDevice_t
	index      uint
	bus        int
	name       string
	uuid       string
}

var nvDevices []Device
var active bool

// Name for specific device
func Name(busNumber int) string {
	for _, d := range nvDevices {
		if d.bus == busNumber {
			return d.name
		}
	}

	return ""
}

// MemoryClock for specific device
func MemoryClock(busNumber int) int {
	for _, d := range nvDevices {
		if d.bus == busNumber {
			return d.MemoryClock()
		}
	}

	return 0
}

// EngineClock for specific device
func EngineClock(busNumber int) int {
	for _, d := range nvDevices {
		if d.bus == busNumber {
			return d.EngineClock()
		}
	}

	return 0
}

// Temperature for specific device
func Temperature(busNumber int) float64 {
	for _, d := range nvDevices {
		if d.bus == busNumber {
			return d.Temperature()
		}
	}

	return 0
}

// FanPercent for specific device
func FanPercent(busNumber int) float64 {
	for _, d := range nvDevices {
		if d.bus == busNumber {
			return d.FanPercent()
		}
	}

	return 0
}

// Active nvml lib
func Active() bool {
	return active
}
