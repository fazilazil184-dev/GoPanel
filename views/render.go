package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func NewTemplateRenderer(templateDir string) (*TemplateRenderer, error) {
	templates := make(map[string]*template.Template)

	pattern := filepath.Join(templateDir, "*.html")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := filepath.Base(file)
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			return nil, err
		}
		templates[name] = tmpl
	}

	return &TemplateRenderer{templates: templates}, nil
}

func (tr *TemplateRenderer) Render(w http.ResponseWriter, name string, data any) error {
	tmpl, exists := tr.templates[name]
	if !exists {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.Execute(w, data)
}
