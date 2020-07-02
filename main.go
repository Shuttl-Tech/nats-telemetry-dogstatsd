package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var (
	Version     string
	Name        = "nats-telemetry-dogstatsd"
	Description = "Export NATS telemetry to Dogstatsd agent"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "dogstatsd.addr",
		EnvVars: []string{"DOGSTATSD_ADDR"},
		Value:   "127.0.0.1:8125",
		Usage:   "Address of the statsd agent with port",
	},
	&cli.StringFlag{
		Name:    "dogstatsd.namespace",
		EnvVars: []string{"DOGSTATSD_NS"},
		Value:   "nats.",
		Usage:   "Name prefix to use for all exposed metrics",
	},
	&cli.StringSliceFlag{
		Name:    "dogstatsd.tags",
		EnvVars: []string{"DOGSTATSD_TAGS"},
		Usage:   "List of tags to export with all metrics",
	},
	&cli.StringFlag{
		Name:     "nats.addr",
		EnvVars:  []string{"NATS_ADDR"},
		Required: true,
		Usage:    "Address of the NATS server",
	},
	&cli.DurationFlag{
		Name:    "metrics.frequency",
		EnvVars: []string{"METRICS_FREQUENCY"},
		Value:   15 * time.Second,
		Usage:   "Frequency at which the metrics will be queries and exported",
	},
}

var app = &cli.App{
	Name:    Name,
	Version: Version,
	Usage:   Description,
	Action:  start,
	Flags:   flags,
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("execution failed. %s", err)
	}
}
