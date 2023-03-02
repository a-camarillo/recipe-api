package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func MuxHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ingredients/", func(w http.ResponseWriter, req *http.Request) {
		
		p := []Ingredient{}
		p = append(p,Ingredient{
			ID: 1,
			Name:"paprika",
		})
		
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
		
	})
	mux.HandleFunc("/recipes/", func(w http.ResponseWriter, req *http.Request) {
		p := []Recipe{}
		p = append(p, Recipe{
			ID: 1,
			Name: "Spicy Curry",
			Ingredients:  []Ingredient{
				Ingredient{
				ID: 1,
				Name: "Paprika",
			},
		}})
		
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)

	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome To The Home Page!")
	})

	return mux
}