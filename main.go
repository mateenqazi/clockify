package main

import (
	"clockify/storage"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	fmt.Println(db)

	/*
		// migration of user table
		err = models.MigrateUser(db)
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
	*/

	app := fiber.New()
	app.Listen(":8080")
}
