package services

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/petrokulybaba/go-basic-framework/configs"
)

var templates = func() *template.Template {
	templates := template.New("")
	err := filepath.Walk(configs.TEMPLATES_DIR, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() {
			template.Must(templates.ParseGlob(strings.TrimRight(path, "/") + "/*" + configs.TEMPLATES_EXTENSION))
		}
		return err
	})
	if err != nil {
		panic(err.Error())
	}
	return templates
}()

func RenderTemplate(w http.ResponseWriter, name string, data ...interface{}) {
	err := templates.ExecuteTemplate(w, name+configs.TEMPLATES_EXTENSION, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
