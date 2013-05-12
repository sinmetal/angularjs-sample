package store

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/store/entry", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Not Found")
		return
	}

	log.Printf("categoryid=%s", r.FormValue("categoryid"))
	log.Printf("itemid=%s", r.FormValue("itemid"))
	log.Printf("name=%s", r.FormValue("name"))
	fmt.Fprintf(w, "OK")
}
