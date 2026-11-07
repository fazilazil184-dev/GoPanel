package views

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func ExistPath(Dir string) (bool, error) {
	info, err := os.Stat(Dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return info.IsDir(), nil
}

func RenderTemplate(w http.ResponseWriter, data any, file string, dir ...string) error {
	var path string

	if len(dir) > 0 && dir[0] != "" {
		isDir, err := ExistPath(dir[0])
		if err != nil {
			return err
		}

		if isDir {
			path = filepath.Join(dir[0], file)
		}
	}

	if path == "" {
		path = file
	}

	_, err := os.Stat(path)

	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return nil
	}

	tpl := template.Must(template.ParseFiles(path))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	return tpl.Execute(w, data)
}
