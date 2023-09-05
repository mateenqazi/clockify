package respository

import (
	"clockify/domain/entity"

	"gorm.io/gorm"
)

type ActivitiesRepository interface {
	GetAllByUserID(userID string) ([]entity.Activities, error)
	Create(activities *entity.Activities) error
	DeleteByID(id string) error
	UpdateName(id int, name string) error
	DuplicateActivity(id string) (*entity.Activities, error)
	SearchActivities(query string, userID string) ([]entity.Activities, error)
}

type ActivitiesDBRepository struct {
	db *gorm.DB
}

func NewActivitiesDBRepository(db *gorm.DB) *ActivitiesDBRepository {
	return &ActivitiesDBRepository{
		db: db,
	}
}

func (r *ActivitiesDBRepository) GetAllByUserID(userID string) ([]entity.Activities, error) {
	var activities []entity.Activities
	if err := r.db.Where("User_id = ?", userID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *ActivitiesDBRepository) Create(activities *entity.Activities) error {
	if err := r.db.Create(activities).Error; err != nil {
		return err
	}
	return nil
}

func (r *ActivitiesDBRepository) DeleteByID(id string) error {
	activities := &entity.Activities{}
	if err := r.db.Where("id = ?", id).Delete(activities).Error; err != nil {
		return err
	}
	return nil
}

func (r *ActivitiesDBRepository) UpdateName(id int, name string) error {
	if err := r.db.Model(&entity.Activities{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

func (r *ActivitiesDBRepository) DuplicateActivity(id string) (*entity.Activities, error) {
	activities := &entity.Activities{}
	newActivities := &entity.Activities{}

	if err := r.db.First(activities, id).Error; err != nil {
		return nil, err
	}

	newActivities.Name = activities.Name + " (copy)"
	newActivities.EndTime = activities.EndTime
	newActivities.StartTime = activities.StartTime
	newActivities.TimeDuration = activities.TimeDuration
	newActivities.UserId = activities.UserId
	newActivities.ProjectId = activities.ProjectId

	if err := r.db.Create(newActivities).Error; err != nil {
		return nil, err
	}

	return newActivities, nil
}

func (r *ActivitiesDBRepository) SearchActivities(query string, userID string) ([]entity.Activities, error) {
	var activities []entity.Activities
	if err := r.db.Model(&activities).Where("name ILIKE ? AND User_id = ?", "%"+query+"%", userID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}
