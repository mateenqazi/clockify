package main

import (
	"clockify/models"
	"clockify/storage"
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

	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	err = models.MigrateProject(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	err = models.MigrateActivities(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	app := fiber.New()
	app.Listen(":8080")
}
