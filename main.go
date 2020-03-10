package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/handlers"
	"github.com/petrokulybaba/go-basic-framework/src/middlewares"
)

func main() {
	// File server
	http.Handle(configs.Routes["assets"]["path"], http.StripPrefix(configs.Routes["assets"]["path"], http.FileServer(http.Dir(configs.ASSETS_DIR))))

	// Handlers
	// All handlers must be wrapped in MakeHandler
	// This gives additional benefits such as checks if a route exist and printing a log
	http.HandleFunc(configs.Routes["index"]["path"], middlewares.MakeHandler(handlers.IndexHandler))
	http.HandleFunc(configs.Routes["login"]["path"], middlewares.MakeHandler(handlers.LoginHandler))

	// Close connection with database
	defer configs.DB.Close()

	fmt.Println("Visit: http://localhost" + configs.PORT)

	log.Fatal(http.ListenAndServe(configs.PORT, nil))
}
