package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/akiko23/golang-tdd/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: sport, kitchen, money, food
---
Body 1
`
		secondBody = `Title: Post 2
Description: Description 2
Tags: programming, IT, python, web
---
S
S
S`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	expectedPosts := []blogposts.Post{
		{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"sport", "kitchen", "money", "food"},
			Body:        "Body 1",
		},
		{
			Title:       "Post 2",
			Description: "Description 2",
			Tags:        []string{"programming", "IT", "python", "web"},
			Body:        "S\nS\nS",
		},
	}

	for current, _ := range posts {
		assertPost(t, posts[current], expectedPosts[current])
	}
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %+v but got %+v", want, got)
	}
}
