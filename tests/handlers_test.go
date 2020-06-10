package handlers_test

import (
	"net/http"
	"testing"

	"github.com/petrokulybaba/go-web-framework/configs"
)

func TestIndexHandler(t *testing.T) {
	response, err := http.Get("http://localhost" + configs.PORT)
	if err != nil {
		t.Fatal(err.Error())
	}
	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong response status code: got %d want %d", status, http.StatusOK)
	}
}
