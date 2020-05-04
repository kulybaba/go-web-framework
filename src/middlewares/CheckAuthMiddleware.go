package middlewares

import (
	"net/http"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/services"
)

func CheckAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(configs.SESSION_COOKIE_NAME)
		if err != nil {
			http.Redirect(w, r, configs.Routes["login"]["path"], http.StatusFound)
			return
		}

		_, err = services.RedisGet(cookie.Value)
		if err != nil {
			http.Redirect(w, r, configs.Routes["login"]["path"], http.StatusFound)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
