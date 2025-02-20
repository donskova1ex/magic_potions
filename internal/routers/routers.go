package routers

import (
	"github.com/donskova1ex/magic_potions/openapi"
	"github.com/gorilla/mux"
)

func NewMuxRouter(
	ingredientController *openapi.IngredientAPIController,
	recipeController *openapi.RecipeAPIController,
	witchController *openapi.WitchAPIController,
) *mux.Router {

	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	apiV1Router.HandleFunc("/witches", witchController.WitchesList).Methods("GET")
	apiV1Router.HandleFunc("/witch/{uuid}", witchController.GetWitchByUUID).Methods("GET")
	apiV1Router.HandleFunc("/witch/{uuid}", witchController.UpdateWitchByUUID).Methods("PUT")
	apiV1Router.HandleFunc("/witch/{uuid}", witchController.DeleteWitchByUUID).Methods("DELETE")

	apiV1Router.HandleFunc("/recipes", recipeController.RecipesList).Methods("GET")
	apiV1Router.HandleFunc("/recipe/{uuid}", recipeController.GetRecipeByUUID).Methods("GET")
	apiV1Router.HandleFunc("/recipe/{uuid}", recipeController.UpdateRecipeByUUID).Methods("PUT")
	apiV1Router.HandleFunc("/recipe/{uuid}", recipeController.DeleteRecipeByUUID).Methods("DELETE")

	apiV1Router.HandleFunc("/ingredients", ingredientController.IngredientsList).Methods("GET")
	apiV1Router.HandleFunc("/ingredient/{uuid}", ingredientController.GetIngredientByUUID).Methods("GET")
	apiV1Router.HandleFunc("/ingredient/{uuid}", ingredientController.UpdateIngredientByUUID).Methods("PUT")
	apiV1Router.HandleFunc("/ingredient/{uuid}", ingredientController.DeleteIngredientByUUID).Methods("DELETE")

	return router
}
