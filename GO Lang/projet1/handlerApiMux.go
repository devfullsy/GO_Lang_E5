package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	dictionary "projet1/directory/directories"

	"github.com/gorilla/mux"
)

func AddHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entry dictionary.Entry
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
			return
		}

		d.Add(entry.Word, entry.Definition)

		w.WriteHeader(http.StatusCreated)
	}
}

func GetHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		word := mux.Vars(r)["word"]

		definition, err := d.Get(word)
		if err != nil {
			http.Error(w, fmt.Sprintf("Mot non trouvé : %s", word), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"word": word, "definition": definition})
	}
}

func RemoveHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		word := mux.Vars(r)["word"]

		d.Remove(word)

		w.WriteHeader(http.StatusOK)
	}
}

func ListHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := d.List()

		w.Write([]byte(result))
	}
}
