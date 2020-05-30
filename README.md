# EdgeTPU Prometheus Exporter

[![Build Status](https://travis-ci.com/adaptant-labs/edgetpu-exporter.svg?branch=master)](https://travis-ci.com/adaptant-labs/edgetpu-exporter#)

This is a Prometheus Exporter for EdgeTPU metrics.

## Building and Installing

Building and installing is the same as for other Go projects:

```
$ go get github.com/adaptant-labs/edgetpu-exporter
```

## Usage

`edgetpu-exporter` can be run as-is without any additional configuration.

```
$ edgetpu-exporter --help
EdgeTPU Prometheus Exporter
Usage: edgetpu-exporter [flags]

  -port int
    	Port to listen to (default 8080)
  -sysfs string
    	Mountpoint of sysfs instance to scan (default "/sys")
```

By default, it will listen on port `8080`, but this can be changed via the `-port` argument.

## Metrics

The following metrics are exported:

| Metric | Description |
|--------|-------------|
| edgetpu_num_devices | Number of EdgeTPU devices |
| edgetpu_temperature_celsius | EdgeTPU device temperature in Celsius (per device) |

Note that not all Apex device types will support temperature checks, there is presently no mechanism by which to obtain
a temperature reading from USB-attached devices.

Viewed from the exporter:

```
# HELP edgetpu_num_devices Number of EdgeTPU devices
# TYPE edgetpu_num_devices gauge
edgetpu_num_devices 1
# HELP edgetpu_temperature_celsius EdgeTPU device temperature in Celsius
# TYPE edgetpu_temperature_celsius gauge
edgetpu_temperature_celsius{name="apex_0"} 49.3
```

## Docker Images

Multi-arch Docker images are available on Docker Hub at [adaptant/edgetpu-exporter].

## Deployment via Kubernetes

`edgetpu-exporter` can be installed directly as a `DaemonSet` on matching nodes:

```
$ kubectl apply -f https://raw.githubusercontent.com/adaptant-labs/edgetpu-exporter/edgetpu-daemonset.yaml
```
 
The node selection criteria depends on the existence of node labels designating the existence of an EdgeTPU within the
node. The list of node labels and their respective labelling sources are listed below:

| Node Label | Labelling Source | Supported Devices |
|------------|------------------|-------------------|
| kkohtaka.org/edgetpu | [EdgeTPU Device Plugin][edgetpu-device-plugin] | USB |
| feature.node.kubernetes.io/usb-fe_1a6e_089a.present | [node-feature-discovery] | USB |
| feature.node.kubernetes.io/pci-0880_1ac1.present | [node-feature-discovery] | PCIe, Coral Dev Board |
| beta.devicetree.org/fsl-imx8mq-phanbell | [k8s-dt-node-labeller] | Coral Dev Board |

## Features and bugs

Please file feature requests and bugs in the [issue tracker][tracker].

## License

`edgetpu-exporter` is licensed under the terms of the Apache 2.0 license, the full
version of which can be found in the LICENSE file included in the distribution.

[tracker]: https://github.com/adaptant-labs/edgetpu-exporter/issues
[adaptant/edgetpu-exporter]: https://hub.docker.com/repository/docker/adaptant/edgetpu-exporter
[k8s-dt-node-labeller]: https://github.com/adaptant-labs/k8s-dt-node-labeller
[node-feature-discovery]: https://github.com/kubernetes-sigs/node-feature/discovery
[edgetpu-device-plugin]: https://github.com/kkohtaka/edgetpu-device-plugin