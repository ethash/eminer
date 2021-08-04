package nvml

import "C"

// MemoryClock returns the current memory clock
func (gpu *Device) MemoryClock() int {
	return 0
}

// EngineClock returns the current engine clock
func (gpu *Device) EngineClock() int {
	return 0
}

// Temperature returns the current temperature
func (gpu *Device) Temperature() float64 {
	return float64(0)
}

// FanPercent returns the current fan speed
func (gpu *Device) FanPercent() float64 {
	return float64(0)
}

// Release nvml
func Release() {
}
