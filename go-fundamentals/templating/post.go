package templating

import "strings"

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
