package endpoints

import "github.com/DataDog/datadog-go/statsd"

type Subsz struct {
	NumSubscriptions float64 `json:"num_subscriptions"`
	NumCache         float64 `json:"num_cache"`
	NumInserts       float64 `json:"num_inserts"`
	NumRemoves       float64 `json:"num_removes"`
	NumMatches       float64 `json:"num_matches"`
	CacheHitRate     float64 `json:"cache_hit_rate"`
	MaxFanout        float64 `json:"max_fanout"`
	AvgFanout        float64 `json:"avg_fanout"`
}

func (v *Subsz) Export(stats *statsd.Client) {
	_ = stats.Gauge("sub.num_subscriptions", v.NumSubscriptions, nil, 1)
	_ = stats.Gauge("sub.num_cache", v.NumCache, nil, 1)
	_ = stats.Gauge("sub.num_inserts", v.NumInserts, nil, 1)
	_ = stats.Gauge("sub.num_removes", v.NumRemoves, nil, 1)
	_ = stats.Gauge("sub.num_matches", v.NumMatches, nil, 1)
	_ = stats.Gauge("sub.cache_hit_rate", v.CacheHitRate, nil, 1)
	_ = stats.Gauge("sub.max_fanout", v.MaxFanout, nil, 1)
	_ = stats.Gauge("sub.avg_fanout", v.AvgFanout, nil, 1)
}
