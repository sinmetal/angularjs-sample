package store

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/store", handler)
}

type Test struct {
	CategoryId int
	ItemId     int
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
	log.Printf("body=%s", body)

	u := map[string]interface{}{}
	err := json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	log.Printf("categoryid=%s", u["categoryid"])
	log.Printf("itemid=%s", u["itemid"])
	log.Printf("name=%s", u["name"])

	decoder := json.NewDecoder(r.Body)
	var test *Test
	if err = decoder.Decode(&test); err != nil {
		c.Errorf("Error decoding Test: %s", err)
		return
	}

	key := datastore.NewKey(c, "Test", "", 1, nil)
	if _, err := datastore.Put(c, key, &test); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK")
}
