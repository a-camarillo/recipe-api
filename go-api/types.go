package main

type CreateIngredient struct {
	Name string `json:"name"`
}

type DeleteIngredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Recipe struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		Name: name,
	}
}
