package favorite

import (
	"../../process/favorite"
	"appengine/aetest"
	"appengine/datastore"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	json := `{"Email":"sinmetal@example.com", "Nickname":"ヴァリトラ"}`
	b := strings.NewReader(json)
	request, _ := http.NewRequest("POST", "/favorite", b)
	response := httptest.NewRecorder()

	favorite.Process(response, request, c)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}

	favos := []favorite.Favorite{}
	ks, err := datastore.NewQuery("Favorite").Order("Created").GetAll(c, &favos)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(favos); i++ {
		favos[i].Id = ks[i].StringID()
	}
	if len(favos) < 1 {
		t.Fatal(len(favos))
	}
	if favos[0].Id != "sinmetal@example.com-_-ヴァリトラ" {
		t.Fatal(favos[0].Id)
	}
}
