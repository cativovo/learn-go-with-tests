package main

import (
	"log"
	"os"
	"path"

	blogposts "github.com/cativovo/learn-go-with-tests/readingfiles"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS(path.Join("testfiles", "posts")))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", posts)
}
