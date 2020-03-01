package middlewares

import (
	"net/http"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/services"
)

// MakeHandler - wrapper for all handlers
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CheckRoute(w, r, fn)
	}
}

// CheckRoute - checks if a route exist
// returns not found page if route does not exist
func CheckRoute(w http.ResponseWriter, r *http.Request, fn func(http.ResponseWriter, *http.Request)) {
	path := r.URL.Path
	for _, v := range configs.Routes {
		if v["path"] == path {
			fn(w, r)
			services.PrintLogOK(path)
			return
		}
	}
	http.NotFound(w, r)
	services.PrintLogNotFound(path)
}
