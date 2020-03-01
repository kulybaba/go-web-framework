package services

import (
	"log"
	"net/http"
)

func PrintLogOK(path string) {
	log.Printf("%d (OK): %s", http.StatusOK, path)
}

func PrintLogNotFound(path string) {
	log.Printf("%d (Page not found): %s", http.StatusNotFound, path)
}
