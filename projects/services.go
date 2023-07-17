package projects

import (
	"clockify/models"
	"errors"
	"fmt"

	"clockify/types"

	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (s *ProjectService) GetAllProject() ([]models.User, error) {
	var users []models.User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *ProjectService) CreateNewProject(proj types.Project) (bool, error) {
	projects := types.Project{
		Name:       proj.Name,
		ClientName: proj.ClientName,
		UserId:     proj.UserId,
	}

	if proj.Name == "" || proj.ClientName == "" {
		return false, errors.New("empty field are not allowed")
	}

	if !(proj.UserId > 0) {
		panic("please login and then create project")
	}

	result := s.db.Create(&projects)
	if result.Error != nil {
		panic("Failed to save data into the database!")
	}
	fmt.Println("Project saved successfully!")

	return true, nil
}

func (s *ProjectService) SearchProject(searchKeyword string, UserId int) ([]models.Project, error) {
	var project []models.Project

	if err := s.db.Where("name ILIKE ? AND User_id = ?", "%"+searchKeyword+"%", UserId).Find(&project).Error; err != nil {
		return project, errors.New("search failed")
	}
	return project, nil
}

func (s *ProjectService) DeleteProject(name string, UserId int) (bool, error) {
	var project []models.Project

	if err := s.db.Where("User_id = ? AND name = ?", UserId, name).Delete(&project); err.Error != nil {
		return false, errors.New("delete failed")
	}
	return true, nil
}