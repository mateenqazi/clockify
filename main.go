package main

import (
	"clockify/projects"
	"clockify/storage"
	"clockify/users"
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
	// helpers.MigrateTable(db)

	// router configurations
	router := mux.NewRouter()
	users.UserSubrouter(router, db)
	projects.ProjectSubrouter(router, db)

	err1 := http.ListenAndServe(":8080", router)
	if err1 != nil {
		log.Fatalln("There's an error with the server,")
	}
}
