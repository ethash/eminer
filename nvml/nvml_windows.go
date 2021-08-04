package nvml

/*
#cgo CFLAGS: -DWIN32
#cgo LDFLAGS:
#include "nvmlbridge.h"
*/
import "C"

import (
	"errors"
	"log"
	"unsafe"
)

func init() {
	if C.init_nvml() == C.NVML_SUCCESS {
		nvDevices, _ = getAllGPUs()
		active = true
	}
}

// MemoryClock returns the current engine clock
func (gpu *Device) MemoryClock() int {
	var result C.nvmlReturn_t
	var clock C.uint

	result = C.bridge_nvmlDeviceGetClockInfo(gpu.nvmldevice, C.NVML_CLOCK_MEM, &clock)
	if result != C.NVML_SUCCESS {
		return int(0)
	}

	return int(clock)
}

// EngineClock returns the current engine clock
func (gpu *Device) EngineClock() int {
	var result C.nvmlReturn_t
	var clock C.uint

	result = C.bridge_nvmlDeviceGetClockInfo(gpu.nvmldevice, C.NVML_CLOCK_GRAPHICS, &clock)
	if result != C.NVML_SUCCESS {
		return int(0)
	}

	return int(clock)
}

// Temperature returns the current temperature of the card in degrees Celsius
func (gpu *Device) Temperature() float64 {
	var result C.nvmlReturn_t
	var ctemp C.uint

	result = C.bridge_nvmlDeviceGetTemperature(gpu.nvmldevice, C.NVML_TEMPERATURE_GPU, &ctemp)
	if result != C.NVML_SUCCESS {
		return float64(0)
	}

	return float64(ctemp)
}

// FanPercent returns the current fan speed of the device, on devices that
// have fans.
func (gpu *Device) FanPercent() float64 {
	speed, err := gpu.intProperty("FanSpeed")
	if err != nil {
		return float64(0)
	}
	return float64(speed)
}

// newDevice is a contstructor function for Device structs. Given an nvmlDevice_t
// object as input, it populates some static property fields and returns a Device
func newDevice(cdevice C.nvmlDevice_t) (*Device, error) {
	device := Device{
		nvmldevice: cdevice,
	}

	uuid, err := device.UUID()
	if err != nil {
		return nil, errors.New("Cannot retrieve UUID property")
	}
	device.uuid = uuid

	name, err := device.Name()
	if err != nil {
		return nil, errors.New("Cannot retrieve Name property")
	}
	device.name = name

	index, err := device.Index()
	if err != nil {
		return nil, errors.New("Cannot retrieve Index property")
	}
	device.index = index

	device.bus = device.BusID()

	return &device, nil
}

// Index by device
func (gpu *Device) Index() (uint, error) {
	return gpu.intProperty("Index")
}

// BusID for device
func (gpu *Device) BusID() int {
	var result C.nvmlReturn_t
	var cpciInfo C.nvmlPciInfo_t

	result = C.bridge_nvmlDeviceGetPciInfo(gpu.nvmldevice, &cpciInfo)
	if result != C.NVML_SUCCESS {
		return 0
	}

	return int(cpciInfo.bus)
}

type cIntPropFunc struct {
	f C.getintProperty
}

var intpropfunctions = map[string]*cIntPropFunc{
	"Index":    {C.getintProperty(C.bridge_nvmlDeviceGetIndex)},
	"FanSpeed": {C.getintProperty(C.bridge_nvmlDeviceGetFanSpeed)},
}

func (gpu *Device) intProperty(property string) (uint, error) {
	var cuintproperty C.uint

	ipf, ok := intpropfunctions[property]
	if ok == false {
		return 0, errors.New("property not found")
	}

	result := C.bridge_get_int_property(ipf.f, gpu.nvmldevice, &cuintproperty)
	if result != C.EXIT_SUCCESS {
		return 0, errors.New("getintProperty bridge returned error")
	}

	return uint(cuintproperty), nil
}

type cTextPropFunc struct {
	f      C.gettextProperty
	length C.uint
}

var textpropfunctions = map[string]*cTextPropFunc{
	"Name": {C.gettextProperty(C.bridge_nvmlDeviceGetName), C.NVML_DEVICE_NAME_BUFFER_SIZE},
	"UUID": {C.gettextProperty(C.bridge_nvmlDeviceGetUUID), C.NVML_DEVICE_UUID_BUFFER_SIZE},
}

// textProperty takes a propertyname as input and then runs the corresponding
// function in the textpropfunctions map, returning the result as a Go string.
//
// textProperty takes care of allocating (and freeing) the text buffers of
// proper size.
func (gpu *Device) textProperty(property string) (string, error) {
	var propvalue string

	// If there isn't a valid entry for this property in the map, there's no reason
	// to process any further
	tpf, ok := textpropfunctions[property]
	if ok == false {
		return "", errors.New("property not found")
	}

	var buf *C.char = genCStringBuffer(uint(tpf.length))
	defer C.free(unsafe.Pointer(buf))

	result := C.bridge_get_text_property(tpf.f, gpu.nvmldevice, buf, tpf.length)
	if result != C.EXIT_SUCCESS {
		return propvalue, errors.New("gettextProperty bridge returned error")
	}

	propvalue = strndup(buf, uint(tpf.length))
	if len(propvalue) > 0 {
		return propvalue, nil
	}

	return "", errors.New("textProperty returned empty string")
}

// InforomImageVersion returns the global inforom image version
func (gpu *Device) InforomImageVersion() (string, error) {
	return gpu.textProperty("InforomImageVersion")
}

// VbiosVersion returns the VBIOS version of the device
func (gpu *Device) VbiosVersion() (string, error) {
	return gpu.textProperty("VbiosVersion")
}

// Name Return the product name of the device, e.g. "Tesla K40m"
func (gpu *Device) Name() (string, error) {
	return gpu.textProperty("Name")
}

// UUID Return the UUID of the device
func (gpu *Device) UUID() (string, error) {
	return gpu.textProperty("UUID")
}

// Return a proper golang error of representation of the nvmlReturn_t error
func (gpu *Device) Error(cerror C.nvmlReturn_t) error {
	var cerrorstring *C.char

	// No need to process anything further if the nvml call succeeded
	if cerror == C.NVML_SUCCESS {
		return nil
	}

	cerrorstring = C.bridge_nvmlErrorString(cerror)
	if cerrorstring == nil {
		// I'm not sure how this could happen, but it's easy to check for
		return errors.New("Error not found in nvml.h")
	}

	return errors.New(C.GoString(cerrorstring))
}

func nvmlDeviceGetCount() (int, error) {
	var count C.uint

	result := C.bridge_nvmlDeviceGetCount(&count)
	if result != C.NVML_SUCCESS {
		return -1, errors.New("nvmlDeviceGetCount failed")
	}

	return int(count), nil
}

func getAllGPUs() ([]Device, error) {
	var devices []Device
	cdevices, err := getAllDevices()
	if err != nil {
		return devices, err
	}

	for _, cdevice := range cdevices {
		device, err := newDevice(cdevice)
		if err != nil {
			break
		}

		devices = append(devices, *device)
	}

	return devices, nil
}

// getAllDevices returns an array of nvmlDevice_t structs representing all GPU
// devices in the system.
func getAllDevices() ([]C.nvmlDevice_t, error) {
	var devices []C.nvmlDevice_t

	deviceCount, err := nvmlDeviceGetCount()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < deviceCount; i++ {
		var device C.nvmlDevice_t
		result := C.bridge_nvmlDeviceGetHandleByIndex(C.uint(i), &device)
		if result != C.NVML_SUCCESS {
			return devices, errors.New("nvmlDeviceGetHandleByIndex returns error")
		}

		devices = append(devices, device)
	}

	if len(devices) == 0 {
		return devices, errors.New("No devices found")
	}

	return devices, nil
}

// Release nvml
func Release() {
	C.bridge_nvmlShutdown()
}
