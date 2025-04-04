package blockrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	br "github.com/per1Peteia/learn_go_with_tests/17_templates/blogrenderer"
)

func TestRenderer(t *testing.T) {
	var aPost = br.Post{
		Title:       "hello world",
		Body:        "*this* is a **post** and a ![test](https://www.example.com/image.jpg))",
		Description: "this is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := br.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []br.Post{{Title: "hello world 1"}, {Title: "hello world 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var aPost = br.Post{
		Title:       "hello world",
		Body:        "this is a post",
		Description: "this is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := br.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
