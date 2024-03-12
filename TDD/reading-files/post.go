package blogposts

import "io"

// Post represents a post on a blog
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

const (
	titleSeparator       = "Tilte: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	var post = Post{Title: string(postData)[7:]}

	return post, nil
}
