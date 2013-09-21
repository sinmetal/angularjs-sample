package pokemon

import (
	"appengine"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Pokemon struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func Process(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handlePokemons(r)
	if err == nil {
		log.Printf("val=%s", val)
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
		id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		log.Printf("id=%s", id)
		return getPokemons(id)
	}
	return nil, fmt.Errorf("method not implemented")
}

func getPokemons(id int64) ([3]Pokemon, error) {
	pokemons := [3]Pokemon{}

	if id == 1 {
		pokemons[0] = Pokemon{Id: 1, Name: "フシギダネ"}
		pokemons[1] = Pokemon{Id: 152, Name: "チコリータ"}
		pokemons[2] = Pokemon{Id: 387, Name: "ナエトル"}
	} else if id == 2 {
		pokemons[0] = Pokemon{Id: 3, Name: "ヒトカゲ"}
		pokemons[1] = Pokemon{Id: 155, Name: "ヒノアラシ"}
		pokemons[2] = Pokemon{Id: 390, Name: "ヒコザル"}
	} else {
		pokemons[0] = Pokemon{Id: 2, Name: "ゼニガメ"}
		pokemons[1] = Pokemon{Id: 158, Name: "ワニノコ"}
		pokemons[2] = Pokemon{Id: 393, Name: "ポッチャマ"}
	}

	return pokemons, nil
}
