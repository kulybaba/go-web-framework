package middlewares

import (
	"net/http"

	"github.com/petrokulybaba/go-web-framework/configs"
	"github.com/petrokulybaba/go-web-framework/src/services"
)

// CheckRouteMiddleware - wrapper for all handlers
// checks if a route exist
// returns not found page if route does not exist
func CheckRouteMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		for _, v := range configs.Routes {
			if v["path"] == path {
				services.PrintLog(method, http.StatusOK, path)
				handler.ServeHTTP(w, r)
				return
			}
		}
		services.PrintLog(method, http.StatusNotFound, path)
		http.NotFound(w, r)
	})
}
