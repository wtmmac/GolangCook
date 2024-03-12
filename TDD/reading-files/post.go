package blogposts

import (
	"bufio"
	"io"
)

// Post represents a post on a blog
type Post struct {
	Title, Description string
}

const (
	titleSeparator       = "Tilte: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	var post = Post{Title: titleLine[7:], Description: descriptionLine[13:]}

	return post, nil
}
