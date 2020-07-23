package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type EdgeTPUDevice struct {
	name string
	path string
}

type EdgeTPUFinder interface {
	 FindDevices() []EdgeTPUDevice
}

type ApexClassFinder struct {}

func (a ApexClassFinder) FindDevices() []EdgeTPUDevice {
	devices := make([]EdgeTPUDevice, 0)

	// Attempt to lookup Apex devices by device class
	files, err := filepath.Glob(sysfsRoot + "/class/apex/apex_*")
	if err != nil {
		return devices
	}

	for _, v := range files {
		device := EdgeTPUDevice{
			path: v,
			name: filepath.Base(v),
		}

		devices = append(devices, device)
	}

	return devices
}

type UsbDeviceFinder struct {}

func readSysfsFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read device attribute %s: %v", filepath.Base(path), err)
	}

	attrVal := strings.TrimSpace(string(data))
	return attrVal, nil
}

func (u UsbDeviceFinder) FindDevices() []EdgeTPUDevice {
	devices := make([]EdgeTPUDevice, 0)

	files, err := filepath.Glob(sysfsRoot + "/bus/usb/devices/*/idVendor")
	if err != nil {
		return devices
	}

	for _, v := range files {
		vid, err := readSysfsFile(v)
		if err != nil {
			continue
		}

		// Match vendor ID
		if vid != "1a6e" {
			continue
		}

		pid, err := readSysfsFile(filepath.Dir(v) + "/idProduct")
		if err != nil {
			continue
		}

		// Match product ID
		if pid != "089a" {
			continue
		}

		path := filepath.Dir(v)
		device := EdgeTPUDevice{
			path: path,
			name: filepath.Base(path),
		}

		devices = append(devices, device)
	}

	return devices
}

func FindEdgeTPUDevices() []EdgeTPUDevice {
	var finder EdgeTPUFinder

	// Search for the class devices first, as this will already include all of the system devices
	if stat, err := os.Stat(sysfsRoot + "/class/apex"); err == nil && stat.IsDir() {
		finder = ApexClassFinder{}
	} else {
		// Fall back on manual bus enumeration via sysfs for USB device discovery
		finder = UsbDeviceFinder{}
	}

	return finder.FindDevices()
}

func (d EdgeTPUDevice) Temperature() float64 {
	data, err := ioutil.ReadFile(d.path + "/temp")
	if err != nil {
		return 0.0
	}

	tempStr := strings.TrimSpace(string(data))
	temp, _ := strconv.ParseFloat(tempStr, 64)

	return temp / 1000
}
