package main

import (
	"context"
	"flag"
	"github.com/Dima191/route-server/internal/app"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	isDebug    bool
	configPath string
)

func init() {
	flag.BoolVar(&isDebug, "debug", false, "enable debug mode")
	flag.StringVar(&configPath, "config-path", "./config/config.env", "path to config file")
}

func main() {
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	logger := Logger(isDebug)

	a, err := app.New(ctx, configPath, logger)
	if err != nil {
		logger.Error("failed to create app", slog.String("error", err.Error()))
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.Run(); err != nil {
			stop()
			logger.Error("failed to run app", slog.String("error", err.Error()))
		}
	}()

	<-ctx.Done()

	a.Stop()
	wg.Wait()
}

func Logger(isDebug bool) *slog.Logger {
	options := &slog.HandlerOptions{
		AddSource: true,
	}

	switch isDebug {
	case true:
		options.Level = slog.LevelDebug
	default:
		options.Level = slog.LevelError
	}
	return slog.New(slog.NewTextHandler(os.Stdout, options))
}
