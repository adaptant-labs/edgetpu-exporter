package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	namespace = "edgetpu"
)

var (
	labels = []string{"name"}
	sysfsRoot = "/sys"
)

type EdgeTPUCollector struct {
	sync.Mutex
	numDevices  prometheus.Gauge
	temperature *prometheus.GaugeVec
}

func NewEdgeTPUCollector() *EdgeTPUCollector {
	return &EdgeTPUCollector{
		numDevices: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "num_devices",
				Help:      "Number of EdgeTPU devices",
			},
		),
		temperature: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "temperature_celsius",
				Help:      "EdgeTPU device temperature in Celsius",
			},
			labels,
		),
	}
}

func (c *EdgeTPUCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.numDevices.Desc()
	c.temperature.Describe(ch)
}

func (c *EdgeTPUCollector) Collect(ch chan<- prometheus.Metric) {
	// Only allow one collection at a time
	c.Lock()
	defer c.Unlock()

	c.temperature.Reset()

	devices := FindEdgeTPUDevices()

	c.numDevices.Set(float64(len(devices)))
	ch <- c.numDevices

	for i := 0; i < len(devices); i++ {
		device := devices[i]
		temp := device.Temperature()
		// Temperature reading is not supported on all devices, skip the ones we don't know anything about
		if temp > 0.0 {
			c.temperature.WithLabelValues(device.name).Set(temp)
		}
	}

	c.temperature.Collect(ch)
}

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "Port to listen to")
	flag.StringVar(&sysfsRoot, "sysfs", "/sys", "Mountpoint of sysfs instance to scan")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "EdgeTPU Prometheus Exporter\n")
		fmt.Fprintf(os.Stderr, "Usage: edgetpu-exporter [flags]\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	prometheus.MustRegister(NewEdgeTPUCollector())

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s...\n", addr)
	log.Fatalf("ListenAndServe error: %v", http.ListenAndServe(addr, promhttp.Handler()))
}
