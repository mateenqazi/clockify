package main

import (
	"clockify/helpers"
	"clockify/storage"
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
	// userService := users.NewUserService(db)
	// projectService := projects.NewProjectService(db)
	// activitiesService := activities.NewActivitiesService(db)

	// // register user
	// creds := types.Credentials{
	// 	Email:    "mat1@gmail.com",
	// 	Password: "securepassword",
	// }

	// helpers.FormatMessage("Register Service Started")

	// result, err := userService.RegisterUser(creds)
	// if err != nil {
	// 	fmt.Println("Error rise Register User ", err)
	// }

	// fmt.Println(result)

	// helpers.FormatMessage("Register Service Ended")

	// // login user
	// helpers.FormatMessage("Login Service Started")

	// creds = types.Credentials{
	// 	Email:    "mat1@gmail.com",
	// 	Password: "securepassword",
	// }

	// user, err := userService.LoginUser(creds)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("you are login with ", user.Email)

	// helpers.FormatMessage("Login Service Ended")

	// helpers.FormatMessage("Create New Project Service Started")

	// project := types.Project{
	// 	Name:       "Project 9",
	// 	ClientName: "Test Client",
	// 	UserId:     user.ID,
	// }

	// proj, err := projectService.CreateNewProject(project)
	// if err != nil {
	// 	fmt.Println("Error rise Register User ", err)
	// }

	// fmt.Println(proj)

	// helpers.FormatMessage("Create New Project Service Ended")

	// helpers.FormatMessage("Search Project Service Started")

	// searchKeyword := "8"
	// searchResult, err := projectService.SearchProject(searchKeyword, user.ID)
	// if err != nil {
	// 	fmt.Println("Error rise Register User ", err)
	// }

	// fmt.Println(searchResult)

	// helpers.FormatMessage("Search Project Service Ended")

	// helpers.FormatMessage("Create Activities Service Started")
	// act := types.Activities{
	// 	Name:      "Second Activity",
	// 	StartTime: time.Now(),
	// 	EndTime:   time.Now().Add(2 * time.Hour),
	// }
	// act.TimeDuration = act.EndTime.Sub(act.StartTime)
	// act.UserId = user.ID
	// act.ProjectId = 1

	// check, err := activitiesService.CreateActivities(act)
	// if err != nil {
	// 	fmt.Println("Error Rise Create Activities ", err)
	// }

	// fmt.Println(check)
	// helpers.FormatMessage("Create Activities Service Ended")

	// helpers.FormatMessage("Duplicate Activities Service Started")

	// duplicate, err := activitiesService.DuplicateActivity(5)
	// if err != nil {
	// 	fmt.Println("Error Rise Duplicate Activities ", err)
	// }

	// fmt.Println(duplicate)
	// helpers.FormatMessage("Duplicate Activities Service Ended")

	// helpers.FormatMessage("Delete Project Service Started")

	// if len(searchResult) > 0 {
	// 	deleteSuccessfully, err := projectService.DeleteProject(searchResult[0].Name, user.ID)
	// 	if err != nil {
	// 		fmt.Println("Error rise Delete Project ", err)
	// 	}

	// 	fmt.Println(deleteSuccessfully)
	// }

	// helpers.FormatMessage("Delete Project Service Ended")

	// helpers.FormatMessage("Delete Activities Service Started")

	// del, err := activitiesService.DeleteActivity(2)
	// if err != nil {
	// 	fmt.Println("Error Rise Delete Activities ", err)
	// }

	// fmt.Println(del)

	// helpers.FormatMessage("Delete Activities Service Ended")

	// migration
	helpers.MigrateTable(db)

	app := fiber.New()
	app.Listen(":8080")
}
