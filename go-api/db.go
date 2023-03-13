package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	err := s.createRecipesTable()
	if err != nil {
		return err
	}
	err = s.createIngredientsTable()
	if err != nil {
		return err
	}
	err = s.createRecipesIngredientsTable()
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createRecipesIngredientsTable() error {
	query := `create table if not exists recipes_ingredients (
		id serial primary key,
		recipe_id int references recipes(id) on delete cascade,
		recipe_name varchar,
		ingredient_id int references ingredients(id) on delete cascade,
		ingredient_name varchar,
		unique (recipe_name,ingredient_name)
		)`
	_, err := s.db.Exec(query)
	return err
}
func (s *PostgresStore) createRecipesTable() error {
	query := `create table if not exists recipes (
		id serial primary key,
		name varchar unique
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createRecipe(rec *CreateRecipe) error {
	query := `insert into recipes (name) 
		values ($1)
		on conflict do nothing;`
	_, err := s.db.Query(query, rec.Name)
	if err != nil {
		return err
	}
	var (
		recID   int
		recName string
	)
	selectRecipeRow := `select id, name from recipes where name = $1`
	recipeRow := s.db.QueryRow(selectRecipeRow, rec.Name)
	if err = recipeRow.Scan(&recID, &recName); err != nil {
		return err
	}
	for _, ing := range rec.Ingredients {
		query = `insert into ingredients (name)
				  values ($1)
				  on conflict do nothing;`
		_, err = s.db.Query(query, ing.Name)
		if err != nil {
			return err
		}
		var (
			ingID   int
			ingName string
		)
		selectIngredientRow := `select id, name from ingredients where name = $1`
		ingredientRow := s.db.QueryRow(selectIngredientRow, ing.Name)
		if err = ingredientRow.Scan(&ingID, &ingName); err != nil {
			return err
		}
		insertRecipesIngredients := `insert into recipes_ingredients (recipe_id, recipe_name, ingredient_id, ingredient_name)
									values ($1, $2, $3, $4)
									on conflict do nothing`
		_, err = s.db.Query(insertRecipesIngredients, recID, recName, ingID, ingName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgresStore) readRecipe(i string) (*Recipe, error) {
	ingredients := []Ingredient{}
	selectRecipeRow := `select id, name from recipes where id = $1;`
	recipeRow := s.db.QueryRow(selectRecipeRow, i)
	var (
		recID   int
		recName string
	)
	if err := recipeRow.Scan(&recID, &recName); err != nil {
		return nil, err
	}
	selectRecipeIngredients := `select ingredient_id, ingredient_name from recipes_ingredients where recipe_id = $1`
	recipeIngredients, err := s.db.Query(selectRecipeIngredients, i)
	if err != nil {
		return nil, err
	}
	var (
		ingID   int
		ingName string
	)
	defer recipeIngredients.Close()
	for recipeIngredients.Next() {
		recipeIngredients.Scan(&ingID, &ingName)
		ingredients = append(ingredients, Ingredient{
			ID:   ingID,
			Name: ingName,
		})
	}
	p := Recipe{
		ID:          recID,
		Name:        recName,
		Ingredients: ingredients,
	}

	return &p, nil
}

func (s *PostgresStore) readRecipes() ([]Recipe, error) {
	p := []Recipe{}
	query := `select id, name from recipes`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		p = append(p, Recipe{
			ID:   id,
			Name: name,
		})
	}
	return p, nil
}

func (s *PostgresStore) deleteRecipe(id string) (*DeleteRecipe, error) {
	query := `delete from recipes where id = $1`
	row := s.db.QueryRow(query, id)
	var (
		deleteid int
		name     string
	)
	if err := row.Scan(&deleteid, &name); err != nil {
		return nil, err
	}
	rec := &DeleteRecipe{
		ID:   deleteid,
		Name: name,
	}
	return rec, nil
}

func (s *PostgresStore) createIngredientsTable() error {
	query := `CREATE TABLE IF NOT EXISTS ingredients (
		id serial primary key,
		name varchar unique
		)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createIngredient(i *Ingredient) error {
	query := `INSERT INTO ingredients (name)
				VALUES ($1)`
	_, err := s.db.Query(query, i.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) readIngredient(i string) (*Ingredient, error) {
	p := Ingredient{}
	query := `select id, name from ingredients where id = $1;`
	row := s.db.QueryRow(query, i)
	var (
		id   int
		name string
	)
	if err := row.Scan(&id, &name); err != nil {
		return nil, err
	}
	p = Ingredient{
		ID:   id,
		Name: name,
	}
	return &p, nil
}

func (s *PostgresStore) readIngredients() ([]Ingredient, error) {
	p := []Ingredient{}
	query := `select id, name from ingredients;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		p = append(p, Ingredient{
			ID:   id,
			Name: name,
		})
	}
	return p, nil
}

func (s *PostgresStore) deleteIngredient(id string) (*DeleteIngredient, error) {
	ing := DeleteIngredient{}
	query := `DELETE FROM ingredients WHERE id = $1;`
	row := s.db.QueryRow(query, id)
	var (
		deleteid int
		name     string
	)
	if err := row.Scan(&deleteid, &name); err != nil {
		return nil, err
	}
	ing = DeleteIngredient{
		ID:   deleteid,
		Name: name,
	}
	return &ing, nil
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=recipe sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
