package main

import (
	"log/slog"
	"os"
	"os/signal"
	"reddit-monitor/config"
	"reddit-monitor/internal"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	cfg := config.Config{
		Username: "reddit username",
		Password: "reddit user password",
		ClientID: "reddit developer client id",
		Secret: "reddit developer secret",
		ServerPort: 8080,
		Limit: 5,
	}

	credentials := reddit.Credentials{ID: cfg.ClientID, Secret: cfg.Secret, Username: cfg.Username, Password: cfg.Password}

	client, err := reddit.NewClient(credentials)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	cache := internal.NewRedditCache()

	monitor := internal.NewRedditMonitor(cache, client, cfg.Limit, logger)

	go monitor.Run(sigChan)
	logger.Info("monitor is running")

	go internal.RegisterEndpoint(cfg.ServerPort, cache, monitor, sigChan, logger)
	logger.Info("endpoints are ready")

	<- sigChan
	logger.Info("application has ended")
}