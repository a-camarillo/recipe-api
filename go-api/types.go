package main

type Ingredient struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Recipe struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		Name: name,
	}
}