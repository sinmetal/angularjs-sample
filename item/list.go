package item

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/item/list", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "[{\"id\" : \"1\", \"name\" : \"キャベツ\"}]")
}
