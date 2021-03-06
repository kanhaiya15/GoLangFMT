package cmd

// this is cmd/root_cmd.go

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"

	"github.com/joho/godotenv"
	config "github.com/kanhaiya15/GoLangFMT/cfg"
	"github.com/kanhaiya15/GoLangFMT/internal/db/sqldb"
	"github.com/kanhaiya15/GoLangFMT/pkg/global"
	"github.com/kanhaiya15/GoLangFMT/pkg/http"
	"github.com/kanhaiya15/GoLangFMT/pkg/lumber"
	"github.com/spf13/cobra"
)

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use:     "mould",
		Long:    `mould is a golang boilerplate for LambdaTest [rpjects]`,
		Version: global.BINARY_VERSION,
		Run:     run,
	}

	// define flags used for this command
	AttachCLIFlags(&rootCmd)

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// timeout in seconds
	const GracefulTimeout = 5000 * time.Millisecond

	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}

	// Load environment variables from .env if available
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: No .env file found\n")
	}

	cfg, err := config.Load(cmd)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err.Error())
	}
	fmt.Printf("cfg : %+v", cfg)
	// patch logconfig file location with root level log file location
	if cfg.LogFile != "" {
		cfg.LogConfig.FileLocation = filepath.Join(cfg.LogFile, "lt.log")
	}

	// You can also use logrus implementation
	// by using lumber.InstanceLogrusLogger
	err = lumber.NewLogger(cfg.LogConfig, cfg.Verbose, lumber.InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate logger %s", err.Error())
	}
	logger := lumber.GetLogger()

	wg.Add(1)
	// setup http server
	go http.Setup(ctx, cfg, &wg, logger)

	// wg.Add(1)
	// setup scheduler
	// go cron.Setup(ctx, cfg, &wg, logger)
	wg.Add(1)
	go sqldb.Setup(ctx, cfg, &wg, logger)
	// listen for C-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// create channel to mark status of waitgroup
	// this is required to brutally kill application in case of
	// timeout
	done := make(chan struct{})

	// asynchronously wait for all the go routines
	go func() {
		// and wait for all go routines
		wg.Wait()
		logger.Debugf("main: all goroutines have finished.")
		close(done)
	}()

	// wait for signal channel
	select {
	case <-c:
		{
			logger.Debugf("main: received C-c - attempting graceful shutdown ....")
			// tell the goroutines to stop
			logger.Debugf("main: telling goroutines to stop")
			cancel()
			select {
			case <-done:
				logger.Debugf("Go routines exited within timeout")
			case <-time.After(GracefulTimeout):
				logger.Errorf("Graceful timeout exceeded. Brutally killing the application")
			}

		}
	case <-done:
		os.Exit(0)
	}

}
