package endpoints

import (
	"github.com/DataDog/datadog-go/statsd"
	"log"
)

type Emitter interface {
	Gauge(name string, value float64, tags []string, rate float64)
	Count(name string, value int64, tags []string, rate float64)
}

type ddEmitter struct {
	client *statsd.Client
}

func (dd *ddEmitter) Gauge(name string, value float64, tags []string, rate float64) {
	err := dd.client.Gauge(name, value, tags, rate)
	if err != nil {
		log.Printf("failed to write Gauge metrics. %s", err)
	}
}

func (dd *ddEmitter) Count(name string, value int64, tags []string, rate float64) {
	err := dd.client.Count(name, value, tags, rate)
	if err != nil {
		log.Printf("failed to write Count metrics. %s", err)
	}
}

func NewDDEmitter(addr string, options ...statsd.Option) (Emitter, error) {
	client, err := statsd.New(addr, options...)
	if err != nil {
		return nil, err
	}

	return &ddEmitter{
		client: client,
	}, nil
}

type exporter interface {
	Export(emitter Emitter)
}

var Sources = map[string]exporter{
	"/varz":     &Varz{},
	"/connz":    &Connz{},
	"/routez":   &Routez{},
	"/gatewayz": &Gatewayz{},
	"/leafz":    &Leafz{},
	"/subsz":    &Subsz{},
}
