package blockrenderer_test

import (
	"bytes"
	br "github.com/per1Peteia/learn_go_with_tests/17_templates/blogrenderer"
	"testing"
)

func TestRenderer(t *testing.T) {
	var aPost = br.Post{
		Title:       "hello world",
		Body:        "this is a post",
		Description: "this is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := br.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1><p>this is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
