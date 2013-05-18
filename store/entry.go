package store

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/store/entry", handler)
}

type Test struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		post(w, r)
	}

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Not Found")
	return

}

func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key := datastore.NewKey(c, "Test", "", 1, nil)
	var t Test
	if _, err := datastore.Put(c, key, &t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("categoryid=%s", r.FormValue("categoryid"))
	log.Printf("itemid=%s", r.FormValue("itemid"))
	log.Printf("name=%s", r.FormValue("name"))
	fmt.Fprintf(w, "OK")
}
