package favorite

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Favorite struct {
	Id          string    `json:"id"`
	PokemonName string    `json:"pokemonName"`
	Nickname    string    `json:nickname`
	Email       string    `json:email`
	Created     time.Time `json:"created"`
}

func (f *Favorite) key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Favorite", fmt.Sprintf("%s-_-%s"), 0, nil)
}

func (t *Favorite) save(c appengine.Context) (*Favorite, error) {
	k, err := datastore.Put(c, t.key(c), t)
	if err != nil {
		return nil, err
	}
	t.Id = k.StringID()
	return t, nil
}

func decodeFavorite(r io.ReadCloser) (*Favorite, error) {
	defer r.Close()
	var Favorite Favorite
	err := json.NewDecoder(r).Decode(&Favorite)
	return &Favorite, err
}

func getAllFavorites(c appengine.Context) ([]Favorite, error) {
	Favorites := []Favorite{}
	ks, err := datastore.NewQuery("Favorite").Order("Created").GetAll(c, &Favorites)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(Favorites); i++ {
		Favorites[i].Id = ks[i].StringID()
	}
	return Favorites, nil
}

func init() {
	http.HandleFunc("/favorite", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handleFavorites(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("favorite error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleFavorites(c appengine.Context, r *http.Request) (interface{}, error) {
	switch r.Method {
	case "POST":
		favorite, err := decodeFavorite(r.Body)
		if err != nil {
			return nil, err
		}
		return favorite.save(c)
	}
	return nil, fmt.Errorf("method not implemented")
}
