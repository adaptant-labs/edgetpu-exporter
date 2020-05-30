package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type EdgeTPUDevice struct {
	name string
	path string
}

func FindEdgeTPUDevices() []EdgeTPUDevice {
	devices := make([]EdgeTPUDevice, 0)

	// Attempt to lookup Apex devices by device class
	files, err := filepath.Glob("/sys/class/apex/apex_*")
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

func (d EdgeTPUDevice) Temperature() float64 {
	data, err := ioutil.ReadFile(d.path + "/temp")
	if err != nil {
		return 0.0
	}

	tempStr := strings.TrimSpace(string(data))
	temp, _ := strconv.ParseFloat(string(tempStr), 64)

	return temp / 1000
}
