package api

import (
	"appengine"
	"net/http"
	"process/favorite"
	"process/item"
	"process/pokemon"
	"process/store"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	c.Infof("request url = %s", r.URL.Path)
	c.Infof("http header = %s", r.Header)

	switch r.URL.Path {
	default:
		http.Error(w, "not found.", http.StatusNotFound)
	case "/item":
		item.Process(w, r)
	case "favorite":
		favorite.Process(w, r)
	case "pokemon":
		pokemon.Process(w, r)
	case "store":
		store.Process(w, r)
	}
}
