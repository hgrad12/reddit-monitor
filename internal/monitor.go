package internal

import (
	"log/slog"
	"os"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Monitor interface {
	Run()
	Initialize(subreddit string)
}

type RedditMonitor struct {
	client *reddit.Client
	cache  *RedditCache
	pipe   chan Result
	workers map[string]bool
	limit int
	logger *slog.Logger
}

func NewRedditMonitor(cache *RedditCache, client *reddit.Client, limit int, logger *slog.Logger) *RedditMonitor {
	p := make(chan Result)
	return &RedditMonitor{
		client: client,
		cache:  cache,
		pipe:   p,
		workers: make(map[string]bool),
		limit: limit,
		logger: logger,
	}
}

func (m *RedditMonitor) Run(signal chan os.Signal) {
	for {
		select {
		case <- signal:
			return
		default:
			res := <- m.pipe
			m.cache.Set(res.Name, res)
		}
	}
}

func (m *RedditMonitor) Initialize(subreddit string, sigChan chan os.Signal) {
	if _, found := m.workers[subreddit]; found {
		return
	}
	s := NewSubReddit(subreddit, m.pipe, m.client, m.limit, m.logger)

	go s.Read(sigChan)

	m.workers[subreddit] = true
}