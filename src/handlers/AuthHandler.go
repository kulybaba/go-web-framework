package handlers

import (
	"net/http"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/forms"
	"github.com/petrokulybaba/go-basic-framework/src/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		services.RenderTemplate(w, configs.Routes["login"]["name"], nil)
	case "POST":
		login := forms.Login{
			r.FormValue("email"),
			r.FormValue("password"),
		}

		if errors := login.Validate(); errors != nil {
			services.RenderTemplate(w, configs.Routes["login"]["name"], map[string]interface{}{
				"errors": errors,
			})
		} else {
			http.Redirect(w, r, configs.Routes["index"]["path"], http.StatusFound)
		}
	}
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		services.RenderTemplate(w, configs.Routes["registration"]["name"], nil)
	case "POST":
		login := forms.Registration{
			r.FormValue("firstName"),
			r.FormValue("lastName"),
			r.FormValue("email"),
			r.FormValue("password"),
			r.FormValue("repeatPassword"),
		}

		if errors := login.Validate(); errors != nil {
			services.RenderTemplate(w, configs.Routes["registration"]["name"], map[string]interface{}{
				"errors": errors,
			})
		} else {
			http.Redirect(w, r, configs.Routes["index"]["path"], http.StatusFound)
		}
	}
}
