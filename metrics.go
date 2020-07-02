package main

import (
	"encoding/json"
	"github.com/Shuttl-Tech/nats-telemetry-dogstatsd/endpoints"
	"log"
	"net/http"
	"strings"
)

func exportMetrics(emitter endpoints.Emitter, server string) {
	for endpoint, target := range endpoints.Sources {
		query(server, endpoint, target)
		go target.Export(emitter)
	}
}

func query(host, endpoint string, target interface{}) {
	endpoint = strings.TrimRight(host, "/") + endpoint
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Printf("failed to query metrics endpoint %s. %s", endpoint, err)
		return
	}

	defer func() {
		cerr := resp.Body.Close()
		if cerr != nil {
			log.Printf("failed to close response body. %s", cerr)
		}
	}()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Printf("failed to decode response from endpoint %s. %s", endpoint, err)
		return
	}
}
