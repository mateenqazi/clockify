package respository

import (
	"clockify/domain/entity"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]entity.User, error)
	Create(user *entity.User) (entity.User, error)
	Delete(id string) error
	FindByEmail(email string) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Create creates a new user in the database.
func (r *UserRepositoryImpl) GetAllUser() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Model(&users).Find(&users); err.Error != nil {
		return nil, err.Error
	}

	return users, nil
}

// Update updates an existing user in the database.
func (r *UserRepositoryImpl) Create(user *entity.User) (entity.User, error) {
	if err := r.db.Model(&user).Create(&user); err != nil {
		log.Println("Failed to create user:", err.Error)
		return *user, err.Error
	}
	return *user, nil

}

// Delete deletes a user from the database.
func (r *UserRepositoryImpl) Delete(userId string) error {
	var users entity.User
	if err := r.db.Model(&users).Where("id = ?", userId).Delete(&users); err != nil {
		return err.Error
	}
	return nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &user, nil
}
