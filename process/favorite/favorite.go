package favorite

import (
	"appengine"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Form struct {
	Id int32 `json:"id"`
}

type Pokemon struct {
	Id int32
	Name string
}

func decodeForm(r io.ReadCloser) (*Form, error) {
	defer r.Close()
	var form Form
	err := json.NewDecoder(r).Decode(&form)
	return &form, err
}

func init() {
	http.HandleFunc("/favorite", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handlePokemons(r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("pokemon error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePokemons(r *http.Request) (interface{}, error) {
	switch r.Method {
	case "GET":
		form, err := decodeForm(r.Body)
		if err != nil {
			return nil, err
		}
		return getPokemons(form)
	}
	return nil, fmt.Errorf("method not implemented")
}


func getPokemons(form *Form) ([3]Pokemon, error) {
	pokemons := [3]Pokemon{}

	charmander := Pokemon{Id: 1, Name: "フシギダネ"}
	pokemons[0] = charmander
	pokemons[1] = charmander
	pokemons[2] = charmander
	return pokemons, nil
}
