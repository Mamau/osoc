package userinfo

import (
	"context"
	"osoc/internal/config"
	"time"

	"osoc/pkg/healthcheck"
	"osoc/pkg/log"
)

type Daemon struct {
	runInterval time.Duration
	stop        chan struct{}
	done        chan struct{}
	status      *healthcheck.StatusVault
	logger      log.Logger
}

func NewDaemon(logger log.Logger, conf config.App) *Daemon {
	return &Daemon{
		runInterval: conf.DaemonRunInterval,
		logger:      logger,
		stop:        make(chan struct{}),
		done:        make(chan struct{}),
		status: healthcheck.NewStatusVault(
			healthcheck.NewProbeStatus("user-daemon", true, "")),
	}
}

func (d *Daemon) Run() {
	defer close(d.done)
	defer d.status.Update(func(ps *healthcheck.ProbeStatus) {
		ps.Ready = false
	})

	tick := time.NewTicker(d.runInterval)
	defer tick.Stop()

	for {
		select {
		case <-d.stop:
			d.logger.Info().Msg("stopping Daemon...")
			return
		case <-tick.C:
			d.logger.Info().Msg("process Daemon...")
			if err := d.process(); err != nil {
				d.logger.Warn().Err(err)
				continue
			}
			d.status.Update(func(ps *healthcheck.ProbeStatus) {
				ps.Ready = true
			})
		}
	}
}

func (d *Daemon) process() error {
	return nil
}

func (d *Daemon) Terminate(ctx context.Context) error {
	close(d.stop)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-d.done:
		return nil
	}
}

func (d *Daemon) Healthcheck(_ context.Context) healthcheck.ProbeStatus {
	return d.status.Load()
}
