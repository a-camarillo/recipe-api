package main

import (
	"net/http"
	"encoding/json"
)

// TODO: Configure DB storage to set up POST, PUT/PATCH?, and DELETE METHODS

func ApiRouteHandler() *http.ServeMux {
	router := NewRouter()

	router.Handle("/ingredients/", IngredientsRouteHandler{})
	router.Handle("/recipes/", RecipesRouteHandler{})

	return router
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	return router
}

type IngredientsRouteHandler struct {}

func (IngredientsRouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		p := []Ingredient{
			Ingredient{
				ID: 1,
				Name: "Spice",
			},
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

type RecipesRouteHandler struct {}

func (RecipesRouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := []Recipe{
			Recipe{
				ID: 1,
				Name: "Spicy Stew",
				Ingredients: []Ingredient{
					Ingredient{
						ID: 1,
						Name: "Spice",
					},
					Ingredient{
						ID: 2,
						Name: "Broth",
					},
				},
			},
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	} 
}

