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
		services.Login(w, r, forms.Login{r.FormValue("email"), r.FormValue("password")})
	}
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		services.RenderTemplate(w, configs.Routes["registration"]["name"], nil)
	case "POST":
		registration := forms.Registration{
			r.FormValue("firstName"),
			r.FormValue("lastName"),
			r.FormValue("email"),
			r.FormValue("password"),
			r.FormValue("repeatPassword"),
		}

		if err := registration.Validate(); err != nil {
			services.RenderTemplate(w, configs.Routes["registration"]["name"], map[string]interface{}{
				"errors": err,
			})
		} else {
			http.Redirect(w, r, configs.Routes["index"]["path"], http.StatusFound)
		}
	}
}
