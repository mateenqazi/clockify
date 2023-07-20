package helpers

import (
	"clockify/models"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
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

func IsPasswordMatch(plainPassword string, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(plainPassword))
}

func HashPassword(password string) (string, error) {
	// Generate the hashed password with a default cost of 10 (higher cost means slower hashing)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func SendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ConvertValueIntoInt(userIdData interface{}) (int, error) {
	switch v := userIdData.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		id, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return id, nil
	default:
		return 0, http.ErrNotSupported // You can use a custom error here if you prefer.
	}

}
