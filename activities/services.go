package activities

import (
	"clockify/models"
	"errors"
	"fmt"

	"clockify/types"

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

func (s *ActivitiesService) GetAllActitives(userId int) ([]models.Activities, error) {
	var activities []models.Activities

	if err := s.db.Where("User_id = ?", userId).Find(&activities).Error; err != nil {
		return activities, errors.New("get all activities failed")
	}

	return activities, nil
}

func (s *ActivitiesService) CreateActivities(act types.Activities) (bool, error) {
	activities := models.Activities{
		Name:         act.Name,
		TimeDuration: act.TimeDuration,
		StartTime:    act.StartTime,
		EndTime:      act.EndTime,
		UserId:       act.UserId,
		ProjectId:    act.ProjectId,
	}

	if act.Name == "" {
		return false, errors.New("empty field are not allowed")
	}

	result := s.db.Create(&activities)
	if result.Error != nil {
		panic("Failed to save data into the database!")
	}
	fmt.Printf("activities %T ", activities)
	fmt.Println("Activities saved successfully!", activities)

	return false, nil
}

func (s *ActivitiesService) DeleteActivity(na int) (bool, error) {
	var activities []models.Activities

	if err := s.db.Delete(activities, na); err.Error != nil {
		return false, errors.New("delete failed")
	}
	fmt.Println("Delete Successfully!")
	return true, nil
}

func (s *ActivitiesService) UpdateName(id int, activityName string) (bool, error) {

	if activityName == "" {
		return false, errors.New("empty name is not allowed")
	}

	if err := s.db.Model(&models.User{}).Where("id = ?", id).Update("name", activityName).Error; err != nil {
		return false, errors.New("error occurred while updating the name")
	}

	fmt.Println("Updated Successfully!")

	return true, nil
}

func (s *ActivitiesService) DuplicateActivity(id int) (bool, error) {
	activities := models.Activities{}
	newActivities := models.Activities{}

	if err := s.db.Where("id = ?", id).First(&activities).Error; err != nil {
		return false, errors.New("duplicate failed")
	}

	newActivities.Name = activities.Name
	newActivities.EndTime = activities.EndTime
	newActivities.StartTime = activities.StartTime
	newActivities.TimeDuration = activities.TimeDuration
	newActivities.UserId = activities.UserId
	newActivities.ProjectId = activities.ProjectId

	result := s.db.Create(&newActivities)

	fmt.Println("Duplicated Successfully!", result)

	return true, nil
}

func (s *ActivitiesService) SearchActivities(searchKeyword string, UserId int) ([]models.Activities, error) {
	var activities []models.Activities

	if err := s.db.Where("name ILIKE ? AND User_id = ?", "%"+searchKeyword+"%", UserId).Find(&activities).Error; err != nil {
		return activities, errors.New("search failed")
	}
	return activities, nil
}
