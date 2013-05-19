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

type Store struct {
	CategoryId int `json:",string"`
	ItemId     int `json:",string"`
	Name       string
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		post(w, r)
	case "GET":
		get(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Not Found")
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	body, _ := ioutil.ReadAll(r.Body)
	var s Store
	if err := json.Unmarshal(body, &s); err != nil {
		c.Errorf("Error unmarshal Store: %s", err)
		return
	}
	c.Infof("CategoryId=%s", s.CategoryId)
	c.Infof("ItemId=%s", s.ItemId)
	c.Infof("Name=%s", s.Name)

	name := fmt.Sprintf("%d-_-%d-_-%s", s.CategoryId, s.ItemId, s.Name)
	key := datastore.NewKey(c, "Store", name, 0, nil)
	if _, err := datastore.Put(c, key, &s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK")
}

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var stores []*Store

	q := datastore.NewQuery("Store")
	_, err := q.GetAll(c, &stores)
	if err != nil {
		c.Errorf("query error : %v", err)
		return
	}

	j, err := json.Marshal(stores)
	if err != nil {
		c.Errorf("json marshal error : %v", err)
		return
	}

	c.Infof("%s", j)
	fmt.Fprintf(w, "%s", j)
}
