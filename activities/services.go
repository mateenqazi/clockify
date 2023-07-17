package activities

import (
	"clockify/helpers"
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
	var activities models.Activities

	if act.Name == "" {
		return false, errors.New("empty field are not allowed")
	}

	fmt.Println(act)

	result := s.db.Create(&activities)
	fmt.Println("Error==>", result.Error)
	if result.Error != nil {
		panic("Failed to save data into the database!")
	}
	fmt.Println("Activities saved successfully!", activities)

	return false, nil
}

func (s *ActivitiesService) LoginUser(creds types.Credentials) (models.User, error) {
	emptyUser := models.User{}
	ok, result, _ := helpers.IsEmailExists(s.db, creds.Email)

	if ok {
		if !helpers.ComparePassword(creds.Password, result.Password) {
			fmt.Println("Password Does not Match")
			fmt.Println("Login Failed!")
			return emptyUser, errors.New("password does not matched")
		}
	}
	fmt.Println("Login Sucessfully!")

	return result, nil
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

	if err := s.db.First(&activities, id); err != nil {
		return false, errors.New("error occurred while updating the name")
	}

	fmt.Println("Updated Successfully!")

	return true, nil
}
