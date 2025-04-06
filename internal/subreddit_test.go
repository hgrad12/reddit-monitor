package internal

import (
	"reflect"
	"testing"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestConvertPostsToResults(t *testing.T) {
	t.Run("return an empty list", func(t *testing.T) {
		t.Parallel()
		var posts []*reddit.Post

		res := ConvertPostsToResults(posts)

		if len(res) != 0 {
			t.Error("list is not empty")
		}
	})

	t.Run("", func(t *testing.T) {
		t.Parallel()
		m := map[string]string{
			"Title1":"Author1",
			"Title2": "Author2",
		}
		posts := []*reddit.Post{
			{
				Title: "Title1",
				Author: "Author1",
			},
			{
				Title: "Title2",
				Author: "Author2",
			},
		}

		res := ConvertPostsToResults(posts)

		if len(res) == 0 {
			t.Error("list is empty")
		}

		for _, post := range posts {
			if _, found := m[post.Title]; !found {
				t.Errorf("list does not contain: %s", post.Title)
			}
		}
	})
}

func TestFindTopPosters(t *testing.T) {
	t.Run("find top users", func(t *testing.T) {
		t.Parallel()
		var users = map[string]int{
			"user1": 3,
			"user2": 1,
			"user3": 4,
		}
		var expected = []TopUsers{
			{
				Author: "user3",
				NumberOfPosts: 4,
			},
			{
				Author: "user1",
				NumberOfPosts: 3,
			},
			{
				Author: "user2",
				NumberOfPosts: 1,
			},
		}

		results := FindTopPosters(users, 3)

		if !reflect.DeepEqual(expected, results) {
			t.Error("lists do not equate")
		}
	})

	t.Run("limit is less than size of users", func(t *testing.T) {
		t.Parallel()
		var users = map[string]int{
			"user1": 3,
			"user2": 1,
			"user3": 4,
		}
		var expected = []TopUsers{
			{
				Author: "user3",
				NumberOfPosts: 4,
			},
			{
				Author: "user1",
				NumberOfPosts: 3,
			},
		}

		results := FindTopPosters(users, 2)

		if !reflect.DeepEqual(expected, results) {
			t.Error("lists do not equate")
		}
	})
}