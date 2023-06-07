package repository

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"recipe_service/src/domain"
	"github.com/google/uuid"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (r *RecipeRepository) GetAllRecipes() ([]domain.Recipe, error) {
	query := "SELECT id, name, description, ingredients, steps FROM recipes"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipes: %w", err)
	}
	defer rows.Close()

	recipes := make([]domain.Recipe, 0)
	for rows.Next() {
		var recipe domain.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, pq.Array(&recipe.Ingredients), pq.Array(&recipe.Steps)); err != nil {
			return nil, fmt.Errorf("failed to scan recipe: %w", err)
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepository) GetRecipeByID(id string) (*domain.Recipe, error) {
	query := "SELECT id, name, description, ingredients, steps FROM recipes WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var recipe domain.Recipe
	if err := row.Scan(&recipe.ID, &recipe.Name, &recipe.Description, pq.Array(&recipe.Ingredients), pq.Array(&recipe.Steps)); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Recipe not found
		}
		return nil, fmt.Errorf("failed to scan recipe: %w", err)
	}

	return &recipe, nil
}

func (r *RecipeRepository) CreateRecipe(recipe *domain.Recipe) error {
	recipeID := uuid.New()


	query := "INSERT INTO recipes (id, name, description, ingredients, steps) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, recipeID, recipe.Name, recipe.Description, pq.Array(recipe.Ingredients), pq.Array(recipe.Steps))
	if err != nil {
		return fmt.Errorf("failed to create recipe: %w", err)
	}
	return nil
}

func (r *RecipeRepository) UpdateRecipe(recipe *domain.Recipe) error {
	query := "UPDATE recipes SET name = $1, description = $2, ingredients = $3, steps = $4 WHERE id = $5"
	_, err := r.db.Exec(query, recipe.Name, recipe.Description, pq.Array(recipe.Ingredients), pq.Array(recipe.Steps), recipe.ID)
	if err != nil {
		return fmt.Errorf("failed to update recipe: %w", err)
	}
	return nil
}

func (r *RecipeRepository) DeleteRecipe(id string) error {
	query := "DELETE FROM recipes WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete recipe: %w", err)
	}
	return nil
}
