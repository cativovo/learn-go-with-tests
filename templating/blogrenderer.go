package templating

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed "templates/*"
var templateFS embed.FS

type templates struct {
	post  *template.Template
	index *template.Template
}

type Renderer struct {
	templates templates
}

func NewRenderer() *Renderer {
	return &Renderer{
		templates: templates{
			post:  template.Must(template.ParseFS(templateFS, "templates/layout.html", "templates/blog.html")),
			index: template.Must(template.ParseFS(templateFS, "templates/layout.html", "templates/index.html")),
		},
	}
}

func (r *Renderer) RenderPost(w io.Writer, post Post) error {
	return r.templates.post.Execute(w, newPostViewModel(post))
}

func (r *Renderer) RenderIndex(w io.Writer, post []Post) error {
	return r.templates.index.Execute(w, post)
}

type postViewModel struct {
	HTMLBody template.HTML
	Post
}

func newPostViewModel(post Post) postViewModel {
	// https://github.com/gomarkdown/markdown/issues/280
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return postViewModel{
		Post:     post,
		HTMLBody: template.HTML(markdown.ToHTML([]byte(post.Body), parser, nil)),
	}
}
