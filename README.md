# EdgeTPU Prometheus Exporter

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
# HELP edgetpu_temperature EdgeTPU device temperature in Celsius
# TYPE edgetpu_temperature gauge
edgetpu_temperature{name="apex_0"} 49.3
```

## Docker Images

Multi-arch Docker images are available on Docker Hub at [adaptant/edgetpu-exporter].

## Features and bugs

Please file feature requests and bugs in the [issue tracker][tracker].

## License

`edgetpu-exporter` is licensed under the terms of the Apache 2.0 license, the full
version of which can be found in the LICENSE file included in the distribution.

[tracker]: https://github.com/adaptant-labs/edgetpu-exporter/issues
[adaptant/edgetpu-exporter]: https://hub.docker.com/repository/docker/adaptant/edgetpu-exporter