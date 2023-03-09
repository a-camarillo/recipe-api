package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func EncodeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func DecodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func HandleFuncCreator(a apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			EncodeJSON(w, http.StatusBadRequest, err)
		}
	}
}

func (s *Server) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/ingredients", HandleFuncCreator(s.IngredientsRootHandler))
	router.HandleFunc("/ingredients/", HandleFuncCreator(s.IngredientsHandler))

	return router
}

// TODO: Configure DB storage to set up POST, PUT/PATCH?, and DELETE METHODS
type Server struct {
	Addr  string
	Store PostgresStore
}

func NewServer(addr string, store PostgresStore) *Server {
	return &Server{
		Addr:  addr,
		Store: store,
	}
}

func (s *Server) Run() {
	router := s.Router()

	http.ListenAndServe(s.Addr, router)
}

func (s *Server) IngredientsRootHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleReadIngredients(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateIngredient(w, r)
	}
	return nil
}

func (s *Server) IngredientsHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" && (strings.TrimPrefix(r.URL.Path, "/ingredients/") == "") {
		return s.handleReadIngredients(w, r)
	}
	if r.Method == "GET" {
		return s.handleReadIngredient(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateIngredient(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteIngredient(w, r)
	}
	return nil
}

func (s *Server) IngredientHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleReadIngredient(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteIngredient(w, r)
	}
	return nil
}

func (s *Server) handleCreateIngredient(w http.ResponseWriter, r *http.Request) error {
	createIngredient := CreateIngredient{}
	if err := json.NewDecoder(r.Body).Decode(&createIngredient); err != nil {
		return err
	}

	ingredient := NewIngredient(createIngredient.Name)
	if err := s.Store.createIngredient(ingredient); err != nil {
		return err
	}

	return EncodeJSON(w, http.StatusCreated, ingredient)
}

func (s *Server) handleReadIngredients(w http.ResponseWriter, r *http.Request) error {
	p, err := s.Store.readIngredients()
	if err != nil {
		log.Fatal(err)
	}
	return EncodeJSON(w, http.StatusOK, p)
}

func (s *Server) handleReadIngredient(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/ingredients/")
	p, err := s.Store.readIngredient(id)
	if err != nil {
		return err
	}
	return EncodeJSON(w, http.StatusOK, p)
}
func (s *Server) handleDeleteIngredient(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/ingredients/")
	ing, err := s.Store.deleteIngredient(id)
	if err != nil {
		return err
	}
	return EncodeJSON(w, http.StatusOK, ing)
}
