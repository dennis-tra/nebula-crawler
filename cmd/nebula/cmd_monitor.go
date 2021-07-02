package main

import (
	"database/sql"
	"strconv"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/dennis-tra/nebula-crawler/pkg/config"
	"github.com/dennis-tra/nebula-crawler/pkg/db"
	"github.com/dennis-tra/nebula-crawler/pkg/metrics"
	"github.com/dennis-tra/nebula-crawler/pkg/monitor"
)

// MonitorCommand contains the receive sub-command configuration.
var MonitorCommand = &cli.Command{
	Name:   "monitor",
	Usage:  "Monitors the network by periodically pinging discovered peers.",
	Action: MonitorAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "dry-run",
			Usage:   "Don't persist results but just ping found peers.",
			EnvVars: []string{"NEBULA_CRAWL_DRY_RUN"},
		},
		&cli.IntFlag{
			Name:        "workers",
			Usage:       "How many concurrent workers should ping peers.",
			EnvVars:     []string{"NEBULA_CRAWL_WORKER_COUNT"},
			DefaultText: strconv.Itoa(config.DefaultConfig.WorkerCount),
			Value:       config.DefaultConfig.WorkerCount,
		},
		&cli.DurationFlag{
			Name:        "dial-timeout",
			Usage:       "How long should be waited before a dial is considered unsuccessful.",
			EnvVars:     []string{"NEBULA_CRAWL_DIAL_TIMEOUT"},
			DefaultText: config.DefaultConfig.DialTimeout.String(),
			Value:       config.DefaultConfig.DialTimeout,
		},
	},
}

// MonitorAction is the function that is called when running pcp receive.
func MonitorAction(c *cli.Context) error {
	log.Infoln("Starting Nebula monitor...")

	// Load configuration file
	ctx, conf, err := config.FillContext(c)
	if err != nil {
		return errors.Wrap(err, "filling context with configuration")
	}
	c.Context = ctx

	// Acquire database handle
	var dbh *sql.DB
	if !c.Bool("dry-run") {
		if dbh, err = db.Open(c.Context); err != nil {
			return err
		}
	}

	// Start prometheus metrics endpoint
	if err = metrics.RegisterListenAndServe(conf.PrometheusHost, conf.PrometheusPort, "monitor"); err != nil {
		return errors.Wrap(err, "initialize metrics")
	}

	// Initialize orchestrator that handles crawling the network.
	m, _ := monitor.NewMonitor(c.Context, dbh)
	go m.StartMonitoring()

	select {
	case <-c.Context.Done():
		// Nebula was asked to stop (e.g. SIGINT) -> tell the monitor to stop
		m.Shutdown()
	case <-m.SigDone():
		// monitor finished autonomously
	}

	return nil
}