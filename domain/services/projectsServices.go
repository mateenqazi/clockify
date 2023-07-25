package services

import (
	"clockify/domain/entity"
	"clockify/domain/respository"
	"clockify/helpers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectServicesInterface interface {
	GetAllProject(w http.ResponseWriter, r *http.Request)
	CreateNewProject(w http.ResponseWriter, r *http.Request)
	SearchProject(w http.ResponseWriter, r *http.Request)
	DeleteProject(w http.ResponseWriter, r *http.Request)
}

type ProjectService struct {
	projectRepository respository.ProjectRepository
}

func NewProjectService(projectRepository respository.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
	}
}

// Implement the functions just as you had them before

func (s *ProjectService) GetAllProject(w http.ResponseWriter, r *http.Request) {
	projects, err := s.projectRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, projects)
}

func (s *ProjectService) CreateNewProject(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := data["name"].(string)
	clientName := data["clientName"].(string)
	userId := data["userId"]

	userIdInt, _ := helpers.ConvertValueIntoInt(userId)

	projects := entity.Project{
		Name:       name,
		ClientName: clientName,
		UserId:     userIdInt,
	}

	if name == "" || clientName == "" {
		http.Error(w, "empty fields are not allowed", http.StatusBadRequest)
		return
	}

	if !(userIdInt > 0) {
		http.Error(w, "userid is missing", http.StatusBadRequest)
		return
	}

	result := s.projectRepository.Create(&projects)
	if result != nil {
		http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
		return
	}

	log.Println("Project saved successfully!")

	helpers.SendJSONResponse(w, http.StatusOK, projects)
}

func (s *ProjectService) SearchProject(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	query := queryParams.Get("query")
	userId := queryParams.Get("userId")

	project, err := s.projectRepository.SearchProject(query, userId)
	if err != nil {
		log.Println("search failed")
		http.Error(w, "search failed", http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, project)
}

func (s *ProjectService) DeleteProject(w http.ResponseWriter, r *http.Request) {

	projectId := mux.Vars(r)["id"]

	if projectId == "" {
		log.Println("delete failed")
		http.Error(w, "project id is empty", http.StatusBadRequest)
		return
	}

	if err := s.projectRepository.Delete(projectId); err != nil {
		log.Println("delete failed")
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	log.Println("project delete sucessfully!")
}
