package favorite

import (
	"../../process/favorite"
	"appengine/aetest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHoge(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	request, _ := http.NewRequest("GET", "/favorite", nil)
	response := httptest.NewRecorder()

	favorite.Process(response, request, c)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}
