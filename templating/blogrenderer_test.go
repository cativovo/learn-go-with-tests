package templating_test

import (
	"bytes"
	"io"
	"testing"

	blogrenderer "github.com/cativovo/learn-go-with-tests/templating"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title: "hello world",
		Body: `# First recipe!
Welcome to my **amazing blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions.`,
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	renderer := blogrenderer.NewRenderer()

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		var buf bytes.Buffer
		if err := renderer.RenderPost(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		// update snapshots
		// UPDATE_SNAPS=true go test ./...
		snaps.MatchSnapshot(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		var buf bytes.Buffer
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := renderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		// update snapshots
		// UPDATE_SNAPS=true go test ./...
		snaps.MatchSnapshot(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	renderer := blogrenderer.NewRenderer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := renderer.RenderPost(io.Discard, aPost); err != nil {
			b.Fatal(err)
		}
	}
}
