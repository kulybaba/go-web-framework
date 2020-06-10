package handlers

import (
	"net/http"

	"github.com/petrokulybaba/go-web-framework/configs"
	"github.com/petrokulybaba/go-web-framework/src/services"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	services.RenderTemplate(w, configs.Routes["index"]["name"], map[string]interface{}{
		"handler":  "IndexHandler",
		"template": "index.html",
	})
}
