package endpoints

import (
	"fmt"
)

type Routez struct {
	ServerID  string   `json:"server_id"`
	Now       string   `json:"now"`
	NumRoutes float64  `json:"num_routes"`
	Routes    []Routes `json:"routes"`
}

type Routes struct {
	Rid           int     `json:"rid"`
	RemoteID      string  `json:"remote_id"`
	DidSolicit    bool    `json:"did_solicit"`
	IP            string  `json:"ip"`
	Port          int     `json:"port"`
	PendingSize   float64 `json:"pending_size"`
	InMsgs        float64 `json:"in_msgs"`
	OutMsgs       float64 `json:"out_msgs"`
	InBytes       float64 `json:"in_bytes"`
	OutBytes      float64 `json:"out_bytes"`
	Subscriptions float64 `json:"subscriptions"`
}

func (v *Routez) Export(stats Emitter) {
	tags := []string{"server_id:" + v.ServerID}

	_ = stats.Gauge("route.num_routes", v.NumRoutes, tags, 1)

	for _, r := range v.Routes {
		rtags := append(tags, fmt.Sprintf("route_id:%d", r.Rid))
		_ = stats.Gauge("route.pending_size", r.PendingSize, rtags, 1)
		_ = stats.Gauge("route.in_msgs", r.InMsgs, rtags, 1)
		_ = stats.Gauge("route.out_msgs", r.OutMsgs, rtags, 1)
		_ = stats.Gauge("route.in_bytes", r.InBytes, rtags, 1)
		_ = stats.Gauge("route.out_bytes", r.OutBytes, rtags, 1)
		_ = stats.Gauge("route.subscriptions", r.Subscriptions, rtags, 1)
	}
}
