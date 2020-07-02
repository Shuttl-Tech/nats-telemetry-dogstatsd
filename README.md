# NATS telemetry dogstatsd

A simple sidecar utility to export NATS telemetry to Dogstatsd.

NATS exposes a lot of telemetry via the [monitoring endpoints][] but does not support any particular metrics sink out of the box. This utility app routinely queries the monitoring endpoints and exports the response data as Dogstatsd metrics.

## Installation

Via `go get`

```shell
go get -u github.com/Shuttl-Tech/nats-telemetry-dogstatsd
```

or from [Github release page][].

You can also pull the docker container from Docker hub.  

Docker images are automatically created from the head of main branch and `latest` tag always points to the head of main. All other tags are immutable and are only created when a new git tag is created.

## Usage

The application requires address to a NATS server. There are more configuration options available, but they are all optional and fall back to sensible defaults.

To start the sidecar in a docker container:

```shell
docker run --rm -it Shuttl-Tech/nats-telemetry-dogstatsd --nats.addr http://server.nats.local:8222
```

or, alternatively the configuration can also be specified using environment variables:

```shell
docker run --rm -it -e NATS_ADDR=http://server.nats.local:8222 Shuttl-Tech/nats-telemetry-dogstatsd
```

For a full list of available configuration parameters please check the help.

## License

MIT License

Copyright (c) 2020 Shuttl

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

[NATS]: https://docs.nats.io/nats-concepts/intro
[monitoring endpoints]: https://docs.nats.io/nats-server/configuration/monitoring#monitoring-endpoints
[Github release page]: https://github.com/Shuttl-Tech/nats-telemetry-dogstatsd/releases
