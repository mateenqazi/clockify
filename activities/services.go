package activities

import (
	"clockify/helpers"
	"clockify/models"
	"clockify/types"
	"encoding/json"
	"errors"
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
	userId := mux.Vars(r)["userid"]

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
	startTime := data["StartTime"]
	endTime := data["EndTime"]
	userId := data["UserId"]
	projectId := data["ProjectId"]

	userIdInt, _ := helpers.ConvertValueIntoInt(userId)
	projectIdInt, _ := helpers.ConvertValueIntoInt(projectId)

	activities := models.Activities{
		Name:         name,
		TimeDuration: timeDuration,
		StartTime:    startTime,
		EndTime:      endTime,
		UserId:       userIdInt,
		ProjectId:    projectIdInt,
	}

	if name == "" {
		return false, errors.New("empty field are not allowed")
	}

	result := s.db.Model(&activities).Create(&activities)
	if result.Error != nil {
		panic("Failed to save data into the database!")
	}

	log.Println("Activities saved successfully!", activities)

	return false, nil

	//dkdkdkkdkdkkdkdkdkdkdkddkkdkdkdkdkkdkdkdkdkdkkdkdkdkd

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

func (s *ActivitiesService) UpdateName(id int, activityName string) (bool, error) {

	if activityName == "" {
		return false, errors.New("empty name is not allowed")
	}

	if err := s.db.Model(&models.Activities{}).Where("id = ?", id).Update("name", activityName).Error; err != nil {
		return false, errors.New("error occurred while updating the name")
	}

	log.Println("Updated Successfully!")

	return true, nil
}

func (s *ActivitiesService) DuplicateActivity(id int) (bool, error) {
	activities := models.Activities{}
	newActivities := models.Activities{}

	if err := s.db.Model(&activities).Where("id = ?", id).First(&activities).Error; err != nil {
		return false, errors.New("duplicate failed")
	}

	newActivities.Name = activities.Name
	newActivities.EndTime = activities.EndTime
	newActivities.StartTime = activities.StartTime
	newActivities.TimeDuration = activities.TimeDuration
	newActivities.UserId = activities.UserId
	newActivities.ProjectId = activities.ProjectId

	result := s.db.Model(&newActivities).Create(&newActivities)
	if result.Error != nil {
		log.Fatal("create failed")
	}

	log.Println("Duplicated Successfully!", result)

	return true, nil
}

func (s *ActivitiesService) SearchActivities(searchKeyword string, UserId int) ([]models.Activities, error) {
	var activities []models.Activities

	if err := s.db.Model(&activities).Where("name ILIKE ? AND User_id = ?", "%"+searchKeyword+"%", UserId).Find(&activities).Error; err != nil {
		return activities, errors.New("search failed")
	}

	return activities, nil
}
