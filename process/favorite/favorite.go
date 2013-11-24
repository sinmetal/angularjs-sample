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
	Id          string    `json:"id" datastore:"-"`
	PokemonName string    `json:"pokemonName"`
	Nickname    string    `json:"nickname"`
	Email       string    `json:"email"`
	Created     time.Time `json:"created"`
}

func Process(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	val, err := handler(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("favorite error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handler(c appengine.Context, r *http.Request) (interface{}, error) {
	switch r.Method {
	case "POST":
		favorite, err := decodeFavorite(r.Body)
		if err != nil {
			return nil, err
		}
		return favorite.save(c)
	case "GET":
		return getAllFavorites(c)
	}

	return nil, fmt.Errorf("method not implemented")
}

func (f *Favorite) key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Favorite", fmt.Sprintf("%s-_-%s", f.Email, f.Nickname), 0, nil)
}

func (f *Favorite) save(c appengine.Context) (*Favorite, error) {
	f.Created = time.Now()
	k, err := datastore.Put(c, f.key(c), f)
	if err != nil {
		return nil, err
	}
	f.Id = k.StringID()
	return f, nil
}

func decodeFavorite(r io.ReadCloser) (*Favorite, error) {
	defer r.Close()
	var Favorite Favorite
	err := json.NewDecoder(r).Decode(&Favorite)
	return &Favorite, err
}

func getAllFavorites(c appengine.Context) ([]Favorite, error) {
	favos := []Favorite{}
	ks, err := datastore.NewQuery("Favorite").Order("Created").GetAll(c, &favos)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(favos); i++ {
		favos[i].Id = ks[i].StringID()
	}
	return favos, nil
}
