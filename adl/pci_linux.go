package adl

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// PCI from sys linux
type PCI struct {
	busNumber int
	dir       string
}

var pciDevices map[int]*PCI

// GetPCIDevice from devices
func GetPCIDevice(busNumber int) (*PCI, error) {
	if p, ok := pciDevices[busNumber]; ok {
		return p, nil
	}

	return nil, errors.New("Device not found")
}

// Temperature from sys
func (p *PCI) Temperature() (temp float64) {
	globs, err := filepath.Glob(filepath.Join(p.dir, "hwmon/*/temp1_input"))
	if err != nil {
		return
	}

	if len(globs) < 1 {
		return
	}

	data, err := ioutil.ReadFile(globs[0])
	if err != nil {
		return
	}

	i, err := strconv.Atoi(string(bytes.Replace(data, []byte("\n"), []byte(""), 1)))
	if err != nil {
		return
	}

	temp = float64(i / 1e3)

	return
}

// FanPercent from sys
func (p *PCI) FanPercent() (percent float64) {
	globs, err := filepath.Glob(filepath.Join(p.dir, "hwmon/*/pwm1"))
	if err != nil {
		return
	}

	if len(globs) < 1 {
		return
	}

	data, err := ioutil.ReadFile(globs[0])
	if err != nil {
		return
	}

	f, err := strconv.Atoi(string(bytes.Replace(data, []byte("\n"), []byte(""), 1)))
	if err != nil {
		return
	}

	max := p.getMaxFanPercent()

	percent = float64(100 * f / max)

	return
}

// FanSetPercent on sys
func (p *PCI) FanSetPercent(fanPercent uint32) (err error) {
	err = checkRootPermission()
	if err != nil {
		return
	}

	var globs []string
	globs, err = filepath.Glob(filepath.Join(p.dir, "hwmon/*"))
	if err != nil {
		return
	}

	if len(globs) < 1 {
		return
	}

	err = ioutil.WriteFile(filepath.Join(globs[0], "pwm1_enable"), []byte("1"), 0664)
	if err != nil {
		return
	}

	max := p.getMaxFanPercent()

	f := max * int(fanPercent) / 100

	err = ioutil.WriteFile(filepath.Join(globs[0], "pwm1"), []byte(strconv.Itoa(f)), 0664)

	return
}

// EngineClock from sys
func (p *PCI) EngineClock() (clock int) {
	data, err := ioutil.ReadFile(filepath.Join(p.dir, "pp_dpm_sclk"))
	if err != nil {
		return
	}

	rows := strings.Split(string(data), "\n")

	for _, row := range rows {
		cols := strings.Split(row, " ")
		if len(cols) == 3 && cols[2] == "*" {
			cls := strings.Replace(cols[1], "Mhz", "", 1)
			clock, _ = strconv.Atoi(cls)
			return
		}
	}

	return
}

// MemoryClock from sys
func (p *PCI) MemoryClock() (clock int) {
	data, err := ioutil.ReadFile(filepath.Join(p.dir, "pp_dpm_mclk"))
	if err != nil {
		return
	}

	rows := strings.Split(string(data), "\n")

	for _, row := range rows {
		cols := strings.Split(row, " ")
		if len(cols) == 3 && cols[2] == "*" {
			cls := strings.Replace(cols[1], "Mhz", "", 1)
			clock, _ = strconv.Atoi(cls)
			return
		}
	}

	return
}

func (p *PCI) getMaxFanPercent() (max int) {
	globs, err := filepath.Glob(filepath.Join(p.dir, "hwmon/*/pwm1_max"))
	if err != nil {
		return
	}

	if len(globs) < 1 {
		return
	}

	data, err := ioutil.ReadFile(globs[0])
	if err != nil {
		return
	}

	max, err = strconv.Atoi(string(bytes.Replace(data, []byte("\n"), []byte(""), 1)))
	if err != nil {
		return
	}

	return
}

func checkRootPermission() (err error) {
	var u *user.User
	u, err = user.Current()
	if err != nil {
		return
	}

	if u.Username != "root" {
		err = errors.New("You need root permission")
		return
	}

	return
}

func getAllPCIDevices() map[int]*PCI {
	devices := make(map[int]*PCI)
	globs, err := filepath.Glob("/sys/bus/pci/devices/*:*:*.0")
	if err != nil {
		return devices
	}

	for _, dir := range globs {
		busByte, err := hex.DecodeString(filepath.Base(dir)[5:7])
		if err != nil || len(busByte) != 1 {
			continue
		}

		busNumber := uint8(busByte[0])
		if busNumber > 0 {
			devices[int(busNumber)] = &PCI{busNumber: int(busNumber), dir: dir}
		}
	}

	return devices
}
