package usecase

import "recipe_service/src/domain"

type RecipeUseCase struct {
	recipeRepository domain.RecipeRepository
}

func NewRecipeUseCase(recipeRepository domain.RecipeRepository) *RecipeUseCase {
	return &RecipeUseCase{
		recipeRepository: recipeRepository,
	}
}

func (uc *RecipeUseCase) GetAllRecipes() ([]domain.Recipe, error) {
	return uc.recipeRepository.GetAllRecipes()
}

func (uc *RecipeUseCase) GetRecipeByID(id string) (*domain.Recipe, error) {
	return uc.recipeRepository.GetRecipeByID(id)
}

func (uc *RecipeUseCase) CreateRecipe(recipe *domain.Recipe) error {
	return uc.recipeRepository.CreateRecipe(recipe)
}

func (uc *RecipeUseCase) UpdateRecipe(recipe *domain.Recipe) error {
	return uc.recipeRepository.UpdateRecipe(recipe)
}

func (uc *RecipeUseCase) DeleteRecipe(id string) error {
	return uc.recipeRepository.DeleteRecipe(id)
}
