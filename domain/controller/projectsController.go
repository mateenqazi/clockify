package controller

import (
	"clockify/domain/respository"
	"clockify/domain/services"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ProjectController struct {
	projectService services.ProjectServicesInterface
}

func NewProjectController(db *gorm.DB) *ProjectController {
	projRepo := respository.NewProjectDBRepository(db)
	projectService := services.NewProjectService(projRepo)
	return &ProjectController{
		projectService: projectService,
	}
}

// RegisterRoutes registers the project-related API routes.
func (c *ProjectController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/projects", c.GetAllProject).Methods(http.MethodGet)
	router.HandleFunc("/projects", c.CreateNewProject).Methods(http.MethodPost)
	router.HandleFunc("/projects/search", c.SearchProject).Methods(http.MethodGet).Queries("query", "{query}", "userId", "{userId}")
	router.HandleFunc("/projects/{id}", c.DeleteProject).Methods(http.MethodDelete)
}

func (c *ProjectController) GetAllProject(w http.ResponseWriter, r *http.Request) {
	c.projectService.GetAllProject(w, r)
}

func (c *ProjectController) CreateNewProject(w http.ResponseWriter, r *http.Request) {
	c.projectService.CreateNewProject(w, r)
}

func (c *ProjectController) SearchProject(w http.ResponseWriter, r *http.Request) {
	c.projectService.SearchProject(w, r)
}

func (c *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	c.projectService.DeleteProject(w, r)
}
