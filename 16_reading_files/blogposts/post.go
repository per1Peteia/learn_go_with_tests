package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	titleLine := scanner.Text()
	scanner.Scan()
	descriptionLine := scanner.Text()

	post := Post{Title: titleLine[7:], Description: descriptionLine[13:]}
	return post, nil
}
