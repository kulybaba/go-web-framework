package services

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/petrokulybaba/go-web-framework/configs"
)

type Template struct {
	Vars map[string]interface{}
}

var templates = func() *template.Template {
	templates := template.New("")
	err := filepath.Walk(configs.TEMPLATES_DIR, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() {
			template.Must(templates.ParseGlob(strings.TrimRight(path, "/") + "/*" + configs.TEMPLATES_EXTENSION))
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return templates
}()

func RenderTemplate(w http.ResponseWriter, name string, vars map[string]interface{}) {
	err := templates.ExecuteTemplate(w, name+configs.TEMPLATES_EXTENSION, Template{vars})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
