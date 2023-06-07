package delivery

import (
	"encoding/json"
	"net/http"

	"recipe_service/src/domain"
	"recipe_service/src/usecase"
)

type RecipeHandler struct {
	recipeUseCase *usecase.RecipeUseCase
}

func NewRecipeHandler(recipeUseCase *usecase.RecipeUseCase) *RecipeHandler {
	return &RecipeHandler{
		recipeUseCase: recipeUseCase,
	}
}

func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request, recipeUseCase *usecase.RecipeUseCase) {
	recipes, err := recipeUseCase.GetAllRecipes()
	if err != nil {
		http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, recipes, http.StatusOK)
}



func (h *RecipeHandler) GetRecipeByID(w http.ResponseWriter, r *http.Request, recipeUseCase *usecase.RecipeUseCase) {
	idStr := r.URL.Path[len("/recipes/"):]

	recipe, err := recipeUseCase.GetRecipeByID(idStr)
	if err != nil {
		http.Error(w, "Failed to fetch recipe", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, recipe, http.StatusOK)
}

func (h *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request, recipeUseCase *usecase.RecipeUseCase) {
	var recipe domain.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = recipeUseCase.CreateRecipe(&recipe)
	if err != nil {
		http.Error(w, "Failed to create recipe", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, recipe, http.StatusCreated)
}



func (h *RecipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request, recipeUseCase *usecase.RecipeUseCase) {
	idStr := r.URL.Path[len("/recipes/"):]

	var recipe domain.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	recipe.ID = idStr
	err = recipeUseCase.UpdateRecipe(&recipe)
	if err != nil {
		http.Error(w, "Failed to update recipe", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, recipe, http.StatusOK)
}


func (h *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request, recipeUseCase *usecase.RecipeUseCase) {
	idStr := r.URL.Path[len("/recipes/"):]

	err := recipeUseCase.DeleteRecipe(idStr)
	if err != nil {
		http.Error(w, "Failed to delete recipe", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
