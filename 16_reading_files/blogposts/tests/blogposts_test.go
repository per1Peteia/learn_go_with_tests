package blogposts_tests

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/per1Peteia/learn_go_with_tests/16_reading_files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 1
Description: Description 1`
	)

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello_world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d, want %d", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
	})

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
