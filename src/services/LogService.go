package services

import "log"

func PrintLog(method string, code int, path string) {
	log.Printf("[%s] %d: %s", method, code, path)
}
