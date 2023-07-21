package projects

import (
	"clockify/helpers"
	"clockify/models"
	"encoding/json"
	"log"
	"net/http"

	"clockify/types"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (s *ProjectService) GetAllProject(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project

	if err := s.db.Model(&projects).Find(&projects).Error; err != nil {
		log.Println(err)
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

	projects := types.Project{
		Name:       name,
		ClientName: clientName,
		UserId:     userIdInt,
	}

	log.Println(projects)

	if name == "" || clientName == "" {
		http.Error(w, "empty fields are not allowed", http.StatusBadRequest)
		return
	}

	if !(userIdInt > 0) {
		http.Error(w, "userid is missing", http.StatusBadRequest)
		return
	}

	result := s.db.Model(&projects).Create(&projects)
	if result.Error != nil {
		http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
		return
	}

	log.Println("Project saved successfully!")

	helpers.SendJSONResponse(w, http.StatusOK, result)
}

func (s *ProjectService) SearchProject(w http.ResponseWriter, r *http.Request) {
	var project []models.Project

	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	query := queryParams.Get("query")
	userid := queryParams.Get("user")

	if err := s.db.Model(&project).Where("name ILIKE ? AND User_id = ?", "%"+query+"%", userid).Find(&project).Error; err != nil {
		log.Println("search failed")
		http.Error(w, "search failed", http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, project)
}

func (s *ProjectService) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var project []models.Project

	projectId := mux.Vars(r)["id"]

	if projectId == "" {
		log.Println("delete failed")
		http.Error(w, "project id is empty", http.StatusBadRequest)
		return
	}

	if err := s.db.Model(&project).Where("id = ?", projectId).Delete(&project); err.Error != nil {
		log.Println("delete failed")
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	log.Println("project delete sucessfully!")

	helpers.SendJSONResponse(w, http.StatusNoContent, nil)
}
