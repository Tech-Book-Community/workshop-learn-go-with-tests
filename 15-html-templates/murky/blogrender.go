package murkyblogrender

import (
	"embed"
	"html/template"
	"io"
)

const (
	postTemplate = "<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags:<ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>"
)

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {

	if err := r.templ.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}

	return nil
}
