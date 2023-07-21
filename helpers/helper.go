package helpers

import (
	"clockify/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"

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

func ConvertValueIntoTimeDuration(value interface{}) (time.Duration, error) {
	switch v := value.(type) {
	case int:
		return time.Duration(v) * time.Second, nil
	case int64:
		return time.Duration(v) * time.Second, nil
	case float32:
		return time.Duration(v*1000) * time.Millisecond, nil
	case float64:
		return time.Duration(v*1000) * time.Millisecond, nil
	default:
		return 0, fmt.Errorf("unsupported data type for conversion to time duration")
	}
}

func ConvertValueToTime(value interface{}) (time.Time, error) {
	switch v := value.(type) {
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case float32:
		return time.Unix(int64(v), 0), nil
	case float64:
		return time.Unix(int64(v), 0), nil
	case string:
		layout := "2006-01-02 15:04:05.999999 -07:00"
		t, err := time.Parse(layout, v)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	default:
		return time.Time{}, fmt.Errorf("unsupported data type for conversion to time.Time")
	}
}
