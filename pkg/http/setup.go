package http

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/LambdaTest/mould/internal/middleware"
	"github.com/LambdaTest/mould/pkg/errs"
	"github.com/LambdaTest/mould/pkg/lumber"
	"github.com/lestrrat-go/backoff"

	"github.com/gin-contrib/cors"

	"github.com/LambdaTest/mould/config"
	global "github.com/LambdaTest/mould/internal/routes/global"
	v1 "github.com/LambdaTest/mould/internal/routes/v1"
	globalConfig "github.com/LambdaTest/mould/pkg/global"
	"github.com/LambdaTest/mould/pkg/ws"
	"github.com/gin-gonic/gin"
)

// Setup initializes all crons on service startup
func StartAPIServer(config *config.Config, ctx context.Context, wg *sync.WaitGroup, logger lumber.Logger) error {
	defer wg.Done()

	// set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	//Initialize logger for packages

	middleware.RegisterLogger(logger)

	logger.Infof("Setting up http handler")
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("authorization", "cache-control", "pragma")
	router.Use(cors.New(corsConfig))

	errChan := make(chan error)

	globalRoute := router.Group("/")
	v1Route := router.Group("/v1.0")

	global.RegisterRoutes(globalRoute, logger)
	ws.RegisterRoutes(globalRoute, logger)
	v1.RegisterRoutes(v1Route, logger)

	// HTTP server instance
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	// channel to signal server process exit
	done := make(chan struct{})
	go func() {
		logger.Infof("Starting server on port %s", config.Port)
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("listen: %#v", err)
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		logger.Infof("Caller has requested graceful shutdown. shutting down the server")
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf("Server Shutdown:", "error", err)
		}
		return nil
	case err := <-errChan:
		return err
	case <-done:
		return nil
	}

}

func Setup(config *config.Config, ctx context.Context, wg *sync.WaitGroup, logger lumber.Logger) error {
	logger.Debugf("Starting API server")

	var policy = backoff.NewExponential(
		backoff.WithInterval(250*time.Millisecond),                         // base interval
		backoff.WithJitterFactor(0.05),                                     // 5% jitter
		backoff.WithMaxRetries(globalConfig.MAX_API_SERVER_START_ATTEMPTS), // If not specified, default number of retries is 10
	)

	b, cancel := policy.Start(context.Background())
	defer cancel()

	for backoff.Continue(b) {
		// channel to mark completion
		done := make(chan struct{})

		// check for context
		select {
		case <-ctx.Done():
			logger.Debugf("Context cancelled. Stopping sink proxy")
			select {
			case <-done:
			case <-time.After(500 * time.Microsecond):
			}
			return nil
		default:
			err := StartAPIServer(config, ctx, wg, logger)
			if err != nil {
				logger.Errorf("API server start error: %s", err)
				logger.Warnf("Restarting api server")
			} else {
				return nil
			}

		}
	}

	// in the case of retry attempt exceed, return with errror
	return errs.ERR_INF_API_MAX_ATTEMPT
}
