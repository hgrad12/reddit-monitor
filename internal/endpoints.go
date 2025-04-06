package internal

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(port int, cache *RedditCache, monitor *RedditMonitor, sigChan chan os.Signal, logger *slog.Logger) {
	a := api{
		cache: cache,
		monitor: monitor,
		sigChan: sigChan,
		logger: logger,
	}
	router := gin.Default()

	router.GET("", a.get)

	router.Run(fmt.Sprintf(":%d", port))
}

type api struct {
	cache *RedditCache
	monitor *RedditMonitor
	sigChan chan os.Signal
	logger *slog.Logger
}

func (a api) get(c *gin.Context) {
	subreddit := c.Query("subreddit")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for _ = range ticker.C {
		res := a.cache.Get(subreddit)

		if res.Name != "" {
			c.JSON(http.StatusOK, res)
		} else {
			a.monitor.Initialize(subreddit, a.sigChan)
		}
	}
}