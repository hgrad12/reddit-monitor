package internal

import (
	"context"
	"log/slog"
	"os"
	"sort"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Reader interface {
	Read()
}

type Subreddit struct {
	subreddit string
	pipe chan Result
	client *reddit.Client
	limit int
	logger *slog.Logger
}

const MaxLimit = 100

func NewSubReddit(subreddit string, pipe chan Result, client *reddit.Client, limit int, logger *slog.Logger) Subreddit {
	return Subreddit{
		subreddit: subreddit,
		pipe: pipe,
		client: client,
		limit: limit,
		logger: logger,
	}
}

func (s Subreddit) Read(sigChan chan os.Signal) {
	for {
		select {
			case <- sigChan:
				return
			default:
				posts, rl, users, err := RetrieveRedditPosts(s.client, s.subreddit, s.limit, s.logger)

			if err != nil {
				s.logger.Error(err.Error())
				continue
			}

			topUsers := FindTopPosters(users, s.limit)

			var topPosts = ConvertPostsToResults(posts)

			s.pipe <- Result{
				Name: s.subreddit, 
				HotPosts: topPosts,
				TopPosters: topUsers,
			}

			//calculate rate limit
			time.Sleep(rl.CalculateRateLimitIntervals())
		}
	}
}

func RetrieveRedditPosts(client *reddit.Client, subreddit string, limit int, logger *slog.Logger) ([]*reddit.Post, RateLimit, map[string]int, error) {
	users := make(map[string]int)
	var p []*reddit.Post
	marker := ""
	counter := 0
	remaining := 0
	var reset time.Time
	for {
		posts, response, err := client.Subreddit.TopPosts(context.Background(), subreddit, &reddit.ListPostOptions{
			ListOptions: reddit.ListOptions{
				Limit: MaxLimit,
				After: marker,
			},
			Time: "all",
		})
		if err != nil {
			logger.Error(err.Error())
			break
		}

		for _, post := range posts {
			users[post.Author]++
		}

		p = append(p, posts...)

		counter += 1
		reset = response.Rate.Reset
		remaining = response.Rate.Remaining

		if len(posts) < MaxLimit {
			break
		}

		marker = posts[len(posts) - 1].FullID
	}

	s := limit
	if len(p) < limit {
		s = len(p)
	}

	rl := RateLimit{
		Count: counter,
		RemainingCalls: remaining,
		Reset: reset,
	}

	return p[:s], rl, users, nil
}


func FindTopPosters(users map[string]int, limit int) []TopUsers {
    var tu []TopUsers
    for author, posts := range users {
        tu = append(tu, TopUsers{author, posts})
    }

    sort.Slice(tu, func(i, j int) bool {
        return tu[i].NumberOfPosts > tu[j].NumberOfPosts
    })

	size := limit

	if len(users) < limit {
		size = len(users)
	}

	return tu[:size]
}

func ConvertPostsToResults(posts []*reddit.Post) []Post {
	var topPosts []Post
	for _, post := range posts {
		topPosts = append(topPosts, Post{
			Title: post.Title,
			Author: post.Author,
		})
	}
	return topPosts
}