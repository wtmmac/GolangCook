package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/wtmmac/GolangCook/TDD/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World
`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M
`
	)

	fs := fstest.MapFS{
		"hello world.md":  {data: []byte(firstBody)},
		"hello world2.md": {data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFS(fs)

	assertNoError(t, err)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
