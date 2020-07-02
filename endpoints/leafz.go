package endpoints

import "github.com/DataDog/datadog-go/statsd"

type Leafz struct {
	ServerID  string  `json:"server_id"`
	Now       string  `json:"now"`
	Leafnodes float64 `json:"leafnodes"`
	Leafs     []Leaf  `json:"leafs"`
}

type Leaf struct {
	Account           string   `json:"account"`
	IP                string   `json:"ip"`
	Port              int      `json:"port"`
	Rtt               string   `json:"rtt"`
	InMsgs            float64  `json:"in_msgs"`
	OutMsgs           float64  `json:"out_msgs"`
	InBytes           float64  `json:"in_bytes"`
	OutBytes          float64  `json:"out_bytes"`
	Subscriptions     float64  `json:"subscriptions"`
	SubscriptionsList []string `json:"subscriptions_list"`
}

func (v *Leafz) Export(stats *statsd.Client) {
	tags := []string{"server_id:" + v.ServerID}
	_ = stats.Gauge("leaf.nodes_count", v.Leafnodes, tags, 1)

	for _, leaf := range v.Leafs {
		leaftags := append(tags, "account:"+leaf.Account)

		_ = stats.Gauge("leaf.in_msgs", leaf.InMsgs, leaftags, 1)
		_ = stats.Gauge("leaf.out_msgs", leaf.OutMsgs, leaftags, 1)
		_ = stats.Gauge("leaf.in_bytes", leaf.InBytes, leaftags, 1)
		_ = stats.Gauge("leaf.out_bytes", leaf.OutBytes, leaftags, 1)
		_ = stats.Gauge("leaf.subscriptions", leaf.Subscriptions, leaftags, 1)
	}
}
