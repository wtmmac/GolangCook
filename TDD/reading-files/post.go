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

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descriptionSeparator):]
	post := Post{Title: title, Description: description}

	return post, nil
}
