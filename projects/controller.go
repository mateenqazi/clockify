package projects

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ProjectSubrouter(r *mux.Router, db *gorm.DB) {

	projectService := NewProjectService(db)

	projectRouter := r.PathPrefix("/projects").Subrouter()

	// apis
	projectRouter.HandleFunc("/create", projectService.CreateNewProject).Methods(http.MethodPost)
	projectRouter.HandleFunc("/delete/{id}", projectService.DeleteProject).Methods(http.MethodDelete)
	projectRouter.HandleFunc("/search", projectService.SearchProject).Methods(http.MethodGet)
	projectRouter.HandleFunc("/", projectService.GetAllProject).Methods(http.MethodGet)
}
