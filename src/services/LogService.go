package services

import (
	"log"
	"net/http"
)

func PrintLogOK(method string, path string) {
	log.Printf("%d %s (OK): %s", http.StatusOK, method, path)
}

func PrintLogNotFound(method string, path string) {
	log.Printf("%d %s (Page not found): %s", http.StatusNotFound, method, path)
}
