package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/handlers"
)

func TestIndexHandler(t *testing.T) {
	request, err := http.NewRequest("GET", configs.Routes["index"]["path"], nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.IndexHandler)
	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
