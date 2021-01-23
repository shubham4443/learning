package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllKeys returns all the keys of the store.
func GetAllKeys(store map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(store)
	}
}

// SetKeys sets new keys in the store.
func SetKeys(store map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(reqBody, &store)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(store)
	}
}

// DeleteAllKeys deletes all the keys of the store.
func DeleteAllKeys(store map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k := range store {
			delete(store, k)
		}
		json.NewEncoder(w).Encode(store)
	}
}

// GetSingleKey returns the specified key in path variable.
func GetSingleKey(store map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		m := make(map[string]interface{})
		m[key] = store[key]
		json.NewEncoder(w).Encode(m)
	}
}

// DeleteSingleKey deletes the specified key in path variable and returns the new store.
func DeleteSingleKey(store map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		delete(store, key)
		json.NewEncoder(w).Encode(store)
	}
}
