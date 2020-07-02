package endpoints

import "github.com/DataDog/datadog-go/statsd"

type exporter interface {
	Export(*statsd.Client)
}

var Sources = map[string]exporter{
	"/varz":     &Varz{},
	"/connz":    &Connz{},
	"/routez":   &Routez{},
	"/gatewayz": &Gatewayz{},
	"/leafz":    &Leafz{},
	"/subsz":    &Subsz{},
}
