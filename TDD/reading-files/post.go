package blogposts

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
