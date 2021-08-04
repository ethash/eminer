package adl

import "C"

// Name of device
func Name(busNumber int) string {
	return ""
}

// FanPercent fetches and returns fan utilization for a device bus number
func FanPercent(busNumber int) float64 {
	return 0
}

// FanSetPercent sets the fan to a percent value for a device bus number
// and returns the ADL return value
func FanSetPercent(busNumber int, fanPercent uint32) error {
	return nil
}

// Temperature fetches and returns temperature for a device bus number
func Temperature(busNumber int) float64 {
	return 0
}

// EngineClock fetches and returns engine clock for a device bus number
func EngineClock(busNumber int) int {
	return 0
}

// EngineSetClock set engine clock for a device bus number
func EngineSetClock(busNumber int, value int) error {
	return nil
}

// MemoryClock fetches and returns memory clock for a device bus number
func MemoryClock(busNumber int) int {
	return 0
}

// MemorySetClock set engine clock for a device bus number
func MemorySetClock(busNumber int, value int) error {
	return nil
}

// Active adl lib
func Active() bool {
	return false
}

// Release ADL
func Release() {}
