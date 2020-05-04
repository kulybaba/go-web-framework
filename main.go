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
	// Close connection with database and Redis
	defer configs.DB.Close()
	defer configs.RedisClient.Close()

	// Secure handlers
	secureHandlers := http.NewServeMux()
	secureHandlers.HandleFunc(configs.Routes["index"]["path"], handlers.IndexHandler)

	// Middlewares for secure handlers
	checkAuthMiddleware := middlewares.CheckAuthMiddleware(secureHandlers)

	// Unsecure handlers
	unsecureHandlers := http.NewServeMux()
	unsecureHandlers.Handle(configs.Routes["index"]["path"], checkAuthMiddleware)
	unsecureHandlers.HandleFunc(configs.Routes["login"]["path"], handlers.LoginHandler)
	unsecureHandlers.HandleFunc(configs.Routes["registration"]["path"], handlers.RegistrationHandler)
	unsecureHandlers.HandleFunc(configs.Routes["logout"]["path"], handlers.LogoutHandler)

	// Middlewares for all handlers
	// All handlers must be wrapped in CheckRouteMiddleware middleware
	// This gives additional benefits such as checks if a route exist and printing a log
	checkRouteMiddleware := middlewares.CheckRouteMiddleware(unsecureHandlers)

	// File server
	fileServer := http.NewServeMux()
	fileServer.Handle(configs.Routes["index"]["path"], checkRouteMiddleware)
	fileServer.Handle(configs.Routes["assets"]["path"], http.StripPrefix(configs.Routes["assets"]["path"], http.FileServer(http.Dir(configs.ASSETS_DIR))))

	fmt.Println("Visit: http://localhost" + configs.PORT)
	log.Fatal(http.ListenAndServe(configs.PORT, fileServer))
}
