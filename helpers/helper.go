package helpers

import (
	"clockify/models"
	"fmt"
	"log"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func IsEmailExists(db *gorm.DB, email string) (bool, models.User, error) {
	users := models.User{}

	result := db.Where("email = ?", email).First(&users)

	if result.Error != nil {
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return false, users, nil
			}
			return false, users, err
		}
	}

	return true, users, nil
}

func MigrateTable(db *gorm.DB) {
	// migration of user table
	err := models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	// migration of project table
	err = models.MigrateProject(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	// migration of activities table
	err = models.MigrateActivities(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
}

func ComparePassword(plainPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(plainPassword))
	return err == nil
}

func HashPassword(password string) string {
	// Generate the hashed password with a default cost of 10 (higher cost means slower hashing)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func FormatMessage(mess string) {
	fmt.Printf("\n\n***********   %v  ***********\n\n", mess)
}
