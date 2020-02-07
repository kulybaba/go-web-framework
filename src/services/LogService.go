package services

import "log"

func PrintLogOK(code int64, path string) {
	log.Printf("%d (OK): %s", code, path)
}

func PrintLogNotFound(code int64, path string) {
	log.Printf("%d (Page not found): %s", code, path)
}
