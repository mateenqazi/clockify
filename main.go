package main

import (
	"clockify/helpers"
	"clockify/storage"
	"clockify/types"
	"clockify/users"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env values
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// database configurations
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// create connections
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	// initialization services
	userService := users.NewUserService(db)

	// register user
	creds := types.Credentials{
		Email:    "mat1@gmail.com",
		Password: "securepassword",
	}

	helpers.FormatMessage("Register Service Started")

	result, err := userService.RegisterUser(creds)
	if err != nil {
		fmt.Println("Error rise Register User ", err)
	}

	fmt.Println(result)

	helpers.FormatMessage("Register Service Ended")

	// login user
	helpers.FormatMessage("Login Service Started")
	creds = types.Credentials{
		Email:    "mat1@gmail.com",
		Password: "securepassword",
	}
	userService.LoginUser(creds)

	helpers.FormatMessage("Login Service Ended")

	// delete user
	// userService.DeleteUser(1)

	// migration
	// helpers.MigrateTable(db)

	app := fiber.New()
	app.Listen(":8080")
}
