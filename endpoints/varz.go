package endpoints

type Varz struct {
	AuthTimeout      int           `json:"auth_timeout"`
	CPU              float64       `json:"cpu"`
	Cluster          *Cluster      `json:"cluster"`
	ConfigLoadTime   string        `json:"config_load_time"`
	Connections      float64       `json:"connections"`
	Cores            int           `json:"cores"`
	Gateway          *Gateway      `json:"gateway"`
	Go               string        `json:"go"`
	Gomaxprocs       int           `json:"gomaxprocs"`
	HTTPHost         string        `json:"http_host"`
	HTTPPort         int           `json:"http_port"`
	HTTPReqStats     *HTTPReqStats `json:"http_req_stats"`
	HTTPSPort        int           `json:"https_port"`
	Host             string        `json:"host"`
	InBytes          int64         `json:"in_bytes"`
	InMsgs           int64         `json:"in_msgs"`
	Leaf             *Leaf         `json:"leaf"`
	Leafnodes        int           `json:"leafnodes"`
	MaxConnections   float64       `json:"max_connections"`
	MaxControlLine   float64       `json:"max_control_line"`
	MaxPayload       float64       `json:"max_payload"`
	MaxPending       float64       `json:"max_pending"`
	Mem              float64       `json:"mem"`
	Now              string        `json:"now"`
	OutBytes         int64         `json:"out_bytes"`
	OutMsgs          int64         `json:"out_msgs"`
	PingInterval     float64       `json:"ping_interval"`
	PingMax          float64       `json:"ping_max"`
	Port             int           `json:"port"`
	Proto            int           `json:"proto"`
	Remotes          int           `json:"remotes"`
	Routes           int           `json:"routes"`
	ServerID         string        `json:"server_id"`
	SlowConsumers    float64       `json:"slow_consumers"`
	Start            string        `json:"start"`
	Subscriptions    float64       `json:"subscriptions"`
	TLSTimeout       float64       `json:"tls_timeout"`
	TotalConnections int64         `json:"total_connections"`
	Uptime           string        `json:"uptime"`
	Version          string        `json:"version"`
	WriteDeadline    int           `json:"write_deadline"`
}

type Cluster struct{}

type HTTPReqStats struct {
	Root     int `json:"/"`
	Connz    int `json:"/connz"`
	Gatewayz int `json:"/gatewayz"`
	Routez   int `json:"/routez"`
	Subsz    int `json:"/subsz"`
	Varz     int `json:"/varz"`
}

func (v *Varz) Export(stats Emitter) {
	tags := []string{"server_id:" + v.ServerID}

	stats.Gauge("server.max_connections", v.MaxConnections, tags, 1)
	stats.Gauge("server.ping_interval", v.PingInterval, tags, 1)
	stats.Gauge("server.ping_max", v.PingMax, tags, 1)
	stats.Gauge("server.max_control_line", v.MaxControlLine, tags, 1)
	stats.Gauge("server.max_payload", v.MaxPayload, tags, 1)
	stats.Gauge("server.max_pending", v.MaxPending, tags, 1)
	stats.Gauge("server.cpu", v.CPU, tags, 1)
	stats.Gauge("server.mem", v.Mem, tags, 1)
	stats.Gauge("server.connections", v.Connections, tags, 1)
	stats.Count("server.total_connections", v.TotalConnections, tags, 1)
	stats.Count("server.in_msg", v.InMsgs, tags, 1)
	stats.Count("server.out_msg", v.OutMsgs, tags, 1)
	stats.Count("server.in_bytes", v.InBytes, tags, 1)
	stats.Count("server.out_bytes", v.OutBytes, tags, 1)
	stats.Gauge("server.subscriptions", v.Subscriptions, tags, 1)
	stats.Gauge("server.slow_consumers", v.SlowConsumers, tags, 1)
}
