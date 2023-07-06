package distribution

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/ximager/ximager/pkg/handlers"
	"github.com/ximager/ximager/pkg/middlewares"
	"github.com/ximager/ximager/pkg/storage"
	"github.com/ximager/ximager/pkg/utils/serializer"
)

func Serve() error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			n := next(c)
			log.Debug().
				Str("method", c.Request().Method).
				Str("path", c.Request().URL.Path).
				Str("query", c.Request().URL.RawQuery).
				Interface("req-header", c.Request().Header).
				Interface("resp-header", c.Response().Header()).
				Int("status", c.Response().Status).
				Msg("Request debugger")
			return n
		}
	}))

	e.Use(middleware.CORS())
	e.Use(middlewares.Healthz())
	e.JSONSerializer = new(serializer.DefaultJSONSerializer)

	if viper.GetInt("log.level") < 1 {
		pprof.Register(e)
	}

	handlers.InitializeDistribution(e)
	err := storage.Initialize()
	if err != nil {
		return err
	}

	go func() {
		log.Info().Str("addr", viper.GetString("http.server")).Msg("Server listening")
		err = e.Start(viper.GetString("http.server"))
		if err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Listening on interface failed")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = e.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Server shutdown failed")
	}

	return nil
}
