package blogposts_tests

import (
	blogposts "github.com/per1Peteia/learn_go_with_tests/16_reading_files/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("Title: Post 1")},
		"hello_world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d, want %d", len(posts), len(fs))
	}
}
