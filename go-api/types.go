package main

type CreateRecipe struct {
	Name        string             `json:"name"`
	Ingredients []CreateIngredient `json:"ingredients"`
}

type DeleteRecipe struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

func NewCreateRecipe(name string, ingredients []CreateIngredient) *CreateRecipe {
	return &CreateRecipe{
		Name:        name,
		Ingredients: ingredients,
	}
}

func NewRecipe(name string) *Recipe {
	return &Recipe{
		Name: name,
	}
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		Name: name,
	}
}
