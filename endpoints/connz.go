package endpoints

import (
	"github.com/DataDog/datadog-go/statsd"
	"time"
)

type Connz struct {
	ServerID       string        `json:"server_id"`
	Now            string        `json:"now"`
	NumConnections int           `json:"num_connections"`
	Total          int           `json:"total"`
	Offset         int           `json:"offset"`
	Limit          int           `json:"limit"`
	Connections    []Connections `json:"connections"`
}

type Connections struct {
	Cid               int      `json:"cid"`
	IP                string   `json:"ip"`
	Port              int      `json:"port"`
	Start             string   `json:"start"`
	LastActivity      string   `json:"last_activity"`
	Rtt               string   `json:"rtt"`
	Uptime            string   `json:"uptime"`
	Idle              string   `json:"idle"`
	PendingBytes      float64  `json:"pending_bytes"`
	InMsgs            float64  `json:"in_msgs"`
	OutMsgs           float64  `json:"out_msgs"`
	InBytes           float64  `json:"in_bytes"`
	OutBytes          float64  `json:"out_bytes"`
	Subscriptions     float64  `json:"subscriptions"`
	Name              string   `json:"name"`
	Lang              string   `json:"lang"`
	Version           string   `json:"version"`
	SubscriptionsList []string `json:"subscriptions_list"`
}

func (v *Connz) Export(stats *statsd.Client) {
	tags := []string{"server_id:" + v.ServerID}

	_ = stats.Gauge("conn.num_total", float64(len(v.Connections)), tags, 1)

	for _, conn := range v.Connections {
		connTags := append(tags, "conn_name:"+conn.Name)

		if d, err := time.ParseDuration(conn.Rtt); err == nil {
			_ = stats.Gauge("conn.rtt_microsec", float64(d.Microseconds()), connTags, 1)
		}

		if d, err := time.ParseDuration(conn.Uptime); err == nil {
			_ = stats.Gauge("conn.uptime_sec", d.Seconds(), connTags, 1)
		}

		if d, err := time.ParseDuration(conn.Idle); err == nil {
			_ = stats.Gauge("conn.idle_sec", d.Seconds(), connTags, 1)
		}

		_ = stats.Gauge("conn.in_msgs", conn.InMsgs, connTags, 1)
		_ = stats.Gauge("conn.out_msgs", conn.OutMsgs, connTags, 1)
		_ = stats.Gauge("conn.in_bytes", conn.InBytes, connTags, 1)
		_ = stats.Gauge("conn.out_bytes", conn.OutBytes, connTags, 1)
		_ = stats.Gauge("conn.pending_bytes", conn.PendingBytes, connTags, 1)
		_ = stats.Gauge("conn.subscriptions", conn.Subscriptions, connTags, 1)
	}
}
