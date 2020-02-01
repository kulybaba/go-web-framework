package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/handlers"
)

func main() {
	// File server
	http.Handle(configs.Routes["assets"]["path"], http.StripPrefix(configs.Routes["assets"]["path"], http.FileServer(http.Dir(configs.ASSETS_DIR))))

	// Handlers
	http.HandleFunc(configs.Routes["index"]["path"], handlers.IndexHandler)

	fmt.Println("Visit: http://localhost" + configs.PORT)

	log.Fatal(http.ListenAndServe(configs.PORT, nil))
}
