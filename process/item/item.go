package item

import (
	"fmt"
	"net/http"
)

func Process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "[{\"id\" : \"1\", \"name\" : \"キャベツ\"}]")
}
