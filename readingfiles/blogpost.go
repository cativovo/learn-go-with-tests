package readingfiles

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, v := range dir {
		post, err := getPost(fileSystem, v.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, filename string) (Post, error) {
	postFile, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	readBody := func() string {
		// skip --
		scanner.Scan()

		var buf bytes.Buffer
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")
	}

	post := Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(),
	}

	return post, nil
}
