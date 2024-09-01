package readingfiles_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/cativovo/learn-go-with-tests/readingfiles"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md": {
			Data: []byte(firstBody),
		},
		"hello world2.md": {
			Data: []byte(secondBody),
		},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	{
		got := len(posts)
		want := len(fs)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	{
		wants := []blogposts.Post{
			{
				Title:       "Post 1",
				Description: "Description 1",
				Tags:        []string{"tdd", "go"},
				Body: `Hello
World`,
			},
			{
				Title:       "Post 2",
				Description: "Description 2",
				Tags:        []string{"rust", "borrow-checker"},
				Body: `B
L
M`,
			},
		}

		assertPost(t, posts, wants)
	}
}

func assertPost(t *testing.T, got, want []blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
