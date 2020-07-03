package endpoints

import (
	"time"
)

type Gatewayz struct {
	ServerID         string               `json:"server_id"`
	Name             string               `json:"name"`
	Now              string               `json:"now"`
	Host             string               `json:"host"`
	Port             int                  `json:"port"`
	OutboundGateways map[string]Gateway   `json:"outbound_gateways"`
	InboundGateways  map[string][]Gateway `json:"inbound_gateways"`
}

type Connection struct {
	Name          string  `json:"name"`
	Cid           int     `json:"cid"`
	IP            string  `json:"ip"`
	Port          int     `json:"port"`
	Start         string  `json:"start"`
	LastActivity  string  `json:"last_activity"`
	Uptime        string  `json:"uptime"`
	Idle          string  `json:"idle"`
	PendingBytes  float64 `json:"pending_bytes"`
	InMsgs        float64 `json:"in_msgs"`
	OutMsgs       float64 `json:"out_msgs"`
	InBytes       float64 `json:"in_bytes"`
	OutBytes      float64 `json:"out_bytes"`
	Subscriptions float64 `json:"subscriptions"`
}

type Gateway struct {
	Configured bool       `json:"configured"`
	Connection Connection `json:"connection"`
}

func (v *Gatewayz) Export(stats Emitter) {
	tags := []string{
		"server_id:" + v.ServerID,
		"name:" + v.Name,
	}

	for name, gw := range v.OutboundGateways {
		dumpGwMetrics(stats, gw, name, "outbound", tags)
	}

	for name, gws := range v.InboundGateways {
		for _, gw := range gws {
			dumpGwMetrics(stats, gw, name, "inbound", tags)
		}
	}
}

func dumpGwMetrics(stats Emitter, gw Gateway, name string, kind string, tags []string) {
	conn := gw.Connection
	gwtags := append(tags, "peer_gw:"+name, "peer_conn:"+conn.Name)
	prefix := "gw." + kind + "."

	if d, err := time.ParseDuration(conn.Uptime); err == nil {
		stats.Gauge(prefix+"conn.uptime_sec", d.Seconds(), gwtags, 1)
	}

	if d, err := time.ParseDuration(conn.Idle); err == nil {
		stats.Gauge(prefix+"conn.idle_sec", d.Seconds(), gwtags, 1)
	}

	stats.Gauge(prefix+"conn.pending_bytes", conn.PendingBytes, gwtags, 1)
	stats.Gauge(prefix+"conn.in_msgs", conn.InMsgs, gwtags, 1)
	stats.Gauge(prefix+"conn.out_msgs", conn.OutMsgs, gwtags, 1)
	stats.Gauge(prefix+"conn.in_bytes", conn.InBytes, gwtags, 1)
	stats.Gauge(prefix+"conn.out_bytes", conn.OutBytes, gwtags, 1)
	stats.Gauge(prefix+"conn.subscriptions", conn.Subscriptions, gwtags, 1)
}
