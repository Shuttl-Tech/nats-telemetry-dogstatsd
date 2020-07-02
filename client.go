package main

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/Shuttl-Tech/nats-telemetry-dogstatsd/endpoints"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	StatsNamespace string
	DogstatsdAddr  string
	DefaultTags    []string
	NATSAddr       string
	Frequency      time.Duration
}

func start(c *cli.Context) error {
	ctx := makeCtx()
	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	metrics, err := statsd.New(config.DogstatsdAddr,
		statsd.WithNamespace(config.StatsNamespace),
		statsd.WithTags(config.DefaultTags),
		statsd.WithoutTelemetry())

	if err != nil {
		return fmt.Errorf("failed to create statsd agent. %s", err)
	}

	return beginExport(ctx, metrics, config.NATSAddr, config.Frequency)
}

func beginExport(ctx context.Context, emitter endpoints.Emitter, server string, frequency time.Duration) error {
	ticker := time.NewTicker(frequency)
	for {
		select {
		case <-ctx.Done():
			return nil

		case <-ticker.C:
			exportMetrics(emitter, server)
		}
	}
}

func parseConfig(c *cli.Context) (*Config, error) {
	config := &Config{
		DogstatsdAddr:  c.String("dogstatsd.addr"),
		StatsNamespace: c.String("dogstatsd.namespace"),
		DefaultTags:    c.StringSlice("dogstatsd.tags"),
		NATSAddr:       c.String("nats.addr"),
		Frequency:      c.Duration("metrics.frequency"),
	}

	return config, nil
}

func makeCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exit
		cancel()
	}()

	return ctx
}
