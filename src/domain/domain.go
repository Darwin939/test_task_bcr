package domain

type Recipe struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

type RecipeRepository interface {
	GetAllRecipes() ([]Recipe, error)
	GetRecipeByID(id string) (*Recipe, error)
	CreateRecipe(recipe *Recipe) error
	UpdateRecipe(recipe *Recipe) error
	DeleteRecipe(id string) error
}
