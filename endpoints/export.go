package endpoints

type Emitter interface {
	Gauge(name string, value float64, tags []string, rate float64) error
	Count(name string, value int64, tags []string, rate float64) error
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
