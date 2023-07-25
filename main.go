package main

import (
	"clockify/domain/controller"
	"clockify/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// database configurations
	config := storage.DataBaseConfig()

	// create connections
	db, err := storage.NewConnection(&config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	// migration
	// helpers.MigrateModels(db)

	// controller initization
	userController := controller.NewUserController(db)
	projectController := controller.NewProjectController(db)
	activitiesController := controller.NewActivitiesController(db)

	// Register the routes for the UserController
	router := mux.NewRouter()
	userController.RegisterRoutes(router)
	projectController.RegisterRoutes(router)
	activitiesController.RegisterRoutes(router)

	err1 := http.ListenAndServe(":8080", router)
	if err1 != nil {
		log.Fatalln("There's an error with the server,")
	}
}
