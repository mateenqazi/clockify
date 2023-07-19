package main

import (
	"clockify/helpers"
	"clockify/projects"
	"clockify/storage"
	"clockify/types"
	"clockify/users"
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
	projectService := projects.NewUserService(db)

	// register user
	creds := types.Credentials{
		Email:    "mat1@gmail.com",
		Password: "securepassword",
	}

	helpers.FormatMessage("Register Service Started")

	result, err := userService.RegisterUser(creds)
	if err != nil {
		log.Fatal("Error rise Register User ", err)
	}

	log.Println(result)

	helpers.FormatMessage("Register Service Ended")

	// login user
	helpers.FormatMessage("Login Service Started")
	creds = types.Credentials{
		Email:    "mat1@gmail.com",
		Password: "securepassword",
	}

	user, err := userService.LoginUser(creds)
	if err != nil {
		panic(err)
	}

	log.Println("you are login with ", user.Email)

	helpers.FormatMessage("Login Service Ended")

	helpers.FormatMessage("Create New Project Service Started")

	project := types.Project{
		Name:       "Project 9",
		ClientName: "Test Client",
		UserId:     user.ID,
	}

	proj, err := projectService.CreateNewProject(project)
	if err != nil {
		log.Fatal("Error rise Register User ", err)
	}

	log.Println(proj)

	helpers.FormatMessage("Create New Project Service Ended")

	helpers.FormatMessage("Search Project Service Started")

	searchKeyword := "8"
	searchResult, err := projectService.SearchProject(searchKeyword, user.ID)
	if err != nil {
		log.Fatal("Error rise Register User ", err)
	}

	log.Println(searchResult)

	helpers.FormatMessage("Search Project Service Ended")

	helpers.FormatMessage("Delete Project Service Started")

	deleteSuccessfully, err := projectService.DeleteProject(searchResult[0].Name, user.ID)
	if err != nil {
		log.Println("Error rise Delete Project ", err)
	}

	log.Println(deleteSuccessfully)
	helpers.FormatMessage("Delete Project Service Ended")

	// migration
	// helpers.MigrateTable(db)

	app := fiber.New()
	app.Listen(":8080")
}
