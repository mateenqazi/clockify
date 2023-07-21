package activities

import (
	"clockify/helpers"
	"clockify/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ActivitiesService struct {
	db *gorm.DB
}

func NewActivitiesService(db *gorm.DB) *ActivitiesService {
	return &ActivitiesService{
		db: db,
	}
}

func (s *ActivitiesService) GetAllActitives(w http.ResponseWriter, r *http.Request) {
	var activities []models.Activities

	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	userId := queryParams.Get("userId")

	if userId == "" {
		log.Println("user id is missing")
		http.Error(w, "user id is missing", http.StatusBadRequest)
		return
	}

	if err := s.db.Model(&activities).Where("User_id = ?", userId).Find(&activities).Error; err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, activities)
}

func (s *ActivitiesService) CreateActivities(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := data["name"].(string)
	timeDuration := data["timeDuration"]
	startTime := data["startTime"]
	endTime := data["endTime"]
	userId := data["userId"]
	projectId := data["projectId"]

	userIdInt, _ := helpers.ConvertValueIntoInt(userId)
	projectIdInt, _ := helpers.ConvertValueIntoInt(projectId)
	timeDurationValue, _ := helpers.ConvertValueIntoTimeDuration(timeDuration)
	startTimeValue, _ := helpers.ConvertValueToTime(startTime)
	endTimeValue, _ := helpers.ConvertValueToTime(endTime)

	activities := models.Activities{
		Name:         name,
		TimeDuration: timeDurationValue,
		StartTime:    startTimeValue,
		EndTime:      endTimeValue,
		UserId:       userIdInt,
		ProjectId:    projectIdInt,
	}

	if name == "" {
		http.Error(w, "empty fields are not allowed", http.StatusBadRequest)
		return
	}

	if !(userIdInt > 0) {
		http.Error(w, "user id is missing", http.StatusBadRequest)
		return
	}

	if !(projectIdInt > 0) {
		http.Error(w, "project id is missing", http.StatusBadRequest)
		return
	}

	result := s.db.Model(&activities).Create(&activities)
	if result.Error != nil {
		http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
		return
	}

	log.Println("Activities saved successfully!")

	helpers.SendJSONResponse(w, http.StatusOK, activities)
}

func (s *ActivitiesService) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	var activities []models.Activities

	activityId := mux.Vars(r)["id"]

	if activityId == "" {
		log.Println("delete failed")
		http.Error(w, "activity id is empty", http.StatusBadRequest)
		return
	}

	if err := s.db.Model(&activities).Delete(&activities, activityId); err.Error != nil {
		log.Println("delete failed")
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	log.Println("Activities Delete Successfully!")

	helpers.SendJSONResponse(w, http.StatusNoContent, nil)
}

func (s *ActivitiesService) UpdateName(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := data["name"].(string)
	id := data["id"]

	if name == "" {
		log.Println("name is missing")
		http.Error(w, "name is missing", http.StatusBadRequest)
		return
	}

	if id == "" {
		log.Println("user id is missing")
		http.Error(w, "user id is missing", http.StatusBadRequest)
		return
	}

	if err := s.db.Model(&models.Activities{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
		return
	}

	log.Println("Updated Successfully!")

	helpers.SendJSONResponse(w, http.StatusOK, "update Sucessfully")
}

func (s *ActivitiesService) DuplicateActivity(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	id := queryParams.Get("id")

	activities := models.Activities{}
	newActivities := models.Activities{}

	if err := s.db.Model(&activities).Where("id = ?", id).First(&activities).Error; err != nil {
		log.Println("duplicated failed")
		http.Error(w, "duplicated failed", http.StatusBadRequest)
		return
	}

	newActivities.Name = activities.Name + " (copy)"
	newActivities.EndTime = activities.EndTime
	newActivities.StartTime = activities.StartTime
	newActivities.TimeDuration = activities.TimeDuration
	newActivities.UserId = activities.UserId
	newActivities.ProjectId = activities.ProjectId

	result := s.db.Model(&newActivities).Create(&newActivities)
	if result.Error != nil {
		log.Println("create failed")
		http.Error(w, "create failed", http.StatusBadRequest)
		return
	}

	log.Println("Duplicated Successfully!", result)

	helpers.SendJSONResponse(w, http.StatusOK, newActivities)
}

func (s *ActivitiesService) SearchActivities(w http.ResponseWriter, r *http.Request) {
	var activities []models.Activities

	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	query := queryParams.Get("query")
	userId := queryParams.Get("userId")

	if err := s.db.Model(&activities).Where("name ILIKE ? AND User_id = ?", "%"+query+"%", userId).Find(&activities).Error; err != nil {
		log.Println("search failed")
		http.Error(w, "search failed", http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, activities)
}
