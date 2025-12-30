package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

type PostRenderer struct {
	templ *template.Template
}

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
