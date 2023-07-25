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

type ActivitiesServicesInterface interface {
	GetAllActitives(w http.ResponseWriter, r *http.Request)
	CreateActivities(w http.ResponseWriter, r *http.Request)
	DeleteActivity(w http.ResponseWriter, r *http.Request)
	UpdateName(w http.ResponseWriter, r *http.Request)
	DuplicateActivity(w http.ResponseWriter, r *http.Request)
	SearchActivities(w http.ResponseWriter, r *http.Request)
}

type ActivitiesService struct {
	activitiesRepository respository.ActivitiesRepository
}

func NewActivitiesService(activitiesRepository respository.ActivitiesRepository) *ActivitiesService {
	return &ActivitiesService{
		activitiesRepository: activitiesRepository,
	}
}

// Implement the functions just as you had them before

func (s *ActivitiesService) GetAllActitives(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	userId := queryParams.Get("userId")

	if userId == "" {
		log.Println("user id is missing")
		http.Error(w, "user id is missing", http.StatusBadRequest)
		return
	}

	activities, err := s.activitiesRepository.GetAllByUserID(userId)
	if err != nil {
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

	activities := entity.Activities{
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

	if err := s.activitiesRepository.Create(&activities); err != nil {
		http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
		return
	}

	log.Println("Activities saved successfully!")

	helpers.SendJSONResponse(w, http.StatusOK, activities)
}

func (s *ActivitiesService) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	activityId := mux.Vars(r)["id"]

	if activityId == "" {
		log.Println("delete failed")
		http.Error(w, "activity id is empty", http.StatusBadRequest)
		return
	}

	if err := s.activitiesRepository.DeleteByID(activityId); err != nil {
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

	userIdInt, _ := helpers.ConvertValueIntoInt(id)

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

	if err := s.activitiesRepository.UpdateName(userIdInt, name); err != nil {
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

	newActivities, err := s.activitiesRepository.DuplicateActivity(id)
	if err != nil {
		log.Println("duplicated failed")
		http.Error(w, "duplicated failed", http.StatusBadRequest)
		return
	}

	log.Println("Duplicated Successfully!", newActivities)

	helpers.SendJSONResponse(w, http.StatusOK, newActivities)
}

func (s *ActivitiesService) SearchActivities(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Access individual query parameters using Get() method
	query := queryParams.Get("query")
	userId := queryParams.Get("userId")

	activities, err := s.activitiesRepository.SearchActivities(query, userId)
	if err != nil {
		log.Println("search failed")
		http.Error(w, "search failed", http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, activities)
}
