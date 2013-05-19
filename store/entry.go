package store

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	http.HandleFunc("/store", handler)
}

type Test struct {
	CategoryId int `json:",string"`
	ItemId     int `json:",string"`
	Name       string
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

	body, _ := ioutil.ReadAll(r.Body)
	var t Test
	if err := json.Unmarshal(body, &t); err != nil {
		c.Errorf("Error unmarshal Test: %s", err)
		return
	}
	c.Infof("CategoryId=%s", t.CategoryId)
	c.Infof("ItemId=%s", t.ItemId)
	c.Infof("Name=%s", t.Name)

	key := datastore.NewKey(c, "Test", "", 1, nil)
	if _, err := datastore.Put(c, key, &t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK")
}
