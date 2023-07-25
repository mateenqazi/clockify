package respository

import (
	"clockify/domain/entity"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAll() ([]entity.Project, error)
	Create(project *entity.Project) error
	Delete(projectId string) error
	SearchProject(query string, userId string) ([]entity.Project, error)
}

type ProjectDBRepository struct {
	db *gorm.DB
}

func NewProjectDBRepository(db *gorm.DB) *ProjectDBRepository {
	return &ProjectDBRepository{
		db: db,
	}
}

func (r *ProjectDBRepository) FindAll() ([]entity.Project, error) {
	var projects []entity.Project
	if err := r.db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectDBRepository) Create(project *entity.Project) error {
	if err := r.db.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProjectDBRepository) Delete(projectId string) error {
	var project entity.Project
	if err := r.db.Model(&project).Where("id = ?", projectId).Delete(&project); err != nil {
		return err.Error
	}
	return nil
}

func (r *ProjectDBRepository) SearchProject(query string, userId string) ([]entity.Project, error) {
	var project []entity.Project

	if err := r.db.Model(&project).Where("name ILIKE ? AND User_id = ?", "%"+query+"%", userId).Find(&project); err != nil {
		return nil, err.Error
	}
	return project, nil
}
