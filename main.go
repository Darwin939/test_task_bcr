package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"recipe_service/src/delivery"
	"recipe_service/src/repository"
	"recipe_service/src/usecase"

	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := waitForDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrateDB(db)

	recipeRepository := repository.NewRecipeRepository(db)
	recipeUseCase := usecase.NewRecipeUseCase(recipeRepository)
	recipeHandler := delivery.NewRecipeHandler(recipeUseCase)

	http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			recipeHandler.GetAllRecipes(w, r, recipeUseCase)
		case http.MethodPost:
			recipeHandler.CreateRecipe(w, r, recipeUseCase)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/recipes/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			recipeHandler.GetRecipeByID(w, r, recipeUseCase)
		case http.MethodPut:
			recipeHandler.UpdateRecipe(w, r, recipeUseCase)
		case http.MethodDelete:
			recipeHandler.DeleteRecipe(w, r, recipeUseCase)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func waitForDatabase() (*sql.DB, error) {
	connectionString := "postgres://user:pass@db_service:5432/db?sslmode=disable"

	for {
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Println("Waiting for the database to be available...")
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Println("Waiting for the database to be available...")
			db.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		return db, nil
	}
}


func migrateDB(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
