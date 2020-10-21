package cron

import (
	"context"
	"sync"

	config "github.com/kanhaiya15/GoLangFMT/cfg"
	"github.com/kanhaiya15/GoLangFMT/pkg/lumber"
	"github.com/robfig/cron/v3"
)

var logger lumber.Logger

// Setup initializes all crons on service startup
func Setup(ctx context.Context, config *config.Config, wg *sync.WaitGroup, initializedLogger lumber.Logger) {
	defer wg.Done()

	logger = initializedLogger
	logger.Infof("Setting up crons")
	c := cron.New()
	c.AddFunc("* * * * *", func() { logger.Infof("Cron alert ...") })
	c.Start()

	select {
	case <-ctx.Done():
		c.Stop()
		logger.Infof("Caller has requested graceful shutdown. Returning.....")
		return
	}

}
