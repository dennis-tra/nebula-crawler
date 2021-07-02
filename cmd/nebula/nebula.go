package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/nebula-crawler/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	_ "net/http/pprof"
)

var (
	// RawVersion and build tag of the
	// PCP command line tool. This is
	// replaced on build via e.g.:
	// -ldflags "-X main.RawVersion=${VERSION}"
	RawVersion  = "dev"
	ShortCommit = "5f3759df" // quake
)

func main() {
	// ShortCommit version tag
	verTag := fmt.Sprintf("v%s+%s", RawVersion, ShortCommit)

	app := &cli.App{
		Name:      "nebula",
		Usage:     "A libp2p DHT crawler that exposes timely information about the network.",
		UsageText: "nebula [global options] command [command options] [arguments...]",
		Authors: []*cli.Author{
			{
				Name:  "Dennis Trautwein",
				Email: "nebula-crawler@dtrautwein.eu",
			},
		},
		Version: verTag,
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				log.SetLevel(log.DebugLevel)
			}
			if c.IsSet("log-level") {
				ll := c.Int("log-level")
				log.SetLevel(log.Level(ll))
				if ll == int(log.TraceLevel) {
					boil.DebugMode = true
				}
			}
			if c.Bool("pprof") {
				go func() {
					if err := http.ListenAndServe("localhost:6060", nil); err != nil {
						log.WithError(err).Warnln("Error serving pprof")
					}
				}()
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Set this flag to enable debug logging.",
				EnvVars: []string{"NEBULA_DEBUG"},
			},
			&cli.IntFlag{
				Name:        "log-level",
				Usage:       "Set this flag to a value from 0 to 6. Overrides the --debug flag.",
				EnvVars:     []string{"NEBULA_LOG_LEVEL"},
				Value:       4,
				DefaultText: "4",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"NEBULA_CONFIG_FILE"},
			},
			//&cli.DurationFlag{
			//	Name:        "min-ping-interval",
			//	Usage:       "The minimum time interval between two consecutive visits of a peer",
			//	EnvVars:     []string{"NEBULA_MIN_PING_INTERVAL"},
			//	DefaultText: config.DefaultConfig.MinPingInterval.String(),
			//	Value:       config.DefaultConfig.MinPingInterval,
			//},
			//&cli.Float64Flag{
			//	Name:        "ping-interval-factor",
			//	Usage:       "The factor with which the next ping timestamp should be calculated",
			//	EnvVars:     []string{"NEBULA_PING_INTERVAL_FACTOR"},
			//	DefaultText: fmt.Sprintf("%f", config.DefaultConfig.PingIntervalFactor),
			//	Value:       config.DefaultConfig.PingIntervalFactor,
			//},
			&cli.IntFlag{
				Name:        "prom-port",
				Usage:       "On which port should prometheus serve the metrics endpoint",
				EnvVars:     []string{"NEBULA_PROMETHEUS_PORT"},
				DefaultText: strconv.Itoa(config.DefaultConfig.PrometheusPort),
				Value:       config.DefaultConfig.PrometheusPort,
			},
			&cli.StringFlag{
				Name:        "prom-host",
				Usage:       "Where should prometheus serve the metrics endpoint",
				EnvVars:     []string{"NEBULA_PROMETHEUS_HOST"},
				DefaultText: config.DefaultConfig.PrometheusHost,
				Value:       config.DefaultConfig.PrometheusHost,
			},
			&cli.BoolFlag{
				Name:    "pprof",
				Usage:   "Enable pprof profiling endpoint on port 6060",
				EnvVars: []string{"NEBULA_PPROF_ENABLED"},
			},
			&cli.StringFlag{
				Name:        "db-host",
				Usage:       "On which host address can nebula reach the database",
				EnvVars:     []string{"NEBULA_DATABASE_HOST"},
				DefaultText: config.DefaultConfig.DatabaseHost,
				Value:       config.DefaultConfig.DatabaseHost,
			},
			&cli.IntFlag{
				Name:        "db-port",
				Usage:       "On which port can nebula reach the database",
				EnvVars:     []string{"NEBULA_DATABASE_PORT"},
				DefaultText: strconv.Itoa(config.DefaultConfig.DatabasePort),
				Value:       config.DefaultConfig.DatabasePort,
			},
			&cli.StringFlag{
				Name:        "db-name",
				Usage:       "The name of the database to use",
				EnvVars:     []string{"NEBULA_DATABASE_NAME"},
				DefaultText: config.DefaultConfig.DatabaseName,
				Value:       config.DefaultConfig.DatabaseName,
			},
			&cli.StringFlag{
				Name:        "db-password",
				Usage:       "The password for the database to use",
				EnvVars:     []string{"NEBULA_DATABASE_PASSWORD"},
				DefaultText: config.DefaultConfig.DatabasePassword,
				Value:       config.DefaultConfig.DatabasePassword,
			},
			&cli.StringFlag{
				Name:        "db-user",
				Usage:       "The user with which to access the database to use",
				EnvVars:     []string{"NEBULA_DATABASE_USER"},
				DefaultText: config.DefaultConfig.DatabaseUser,
				Value:       config.DefaultConfig.DatabaseUser,
			},
		},
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			CrawlCommand,
			MonitorCommand,
		},
	}

	sigs := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		sig := <-sigs
		log.Printf("Received %s - Stopping...\n", sig.String())
		signal.Stop(sigs)
		cancel()
	}()

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}