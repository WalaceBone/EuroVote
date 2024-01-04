package main

import (
	"EuroVote/cmd/models"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Starting the application...")
	fmt.Println(os.Getenv("PORT"))

	var app models.App
	app.Initialize()
	app.Router.Run()
	// Config, err := loadConfig()

	// if err != nil {
	// 	panic(err)
	// }

	// // Connect to the database
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file:", err)
	// }

	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")

	// log.Println(dbHost)
	// log.Println(dbUser)
	// log.Println(dbPassword)
	// log.Println(dbName)
	// log.Println(dbPort)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	// // END: ed8c6549bwf9

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// app := models.App{
	// 	Router: gin.Default(),
	// 	Config: Config,
	// 	DB:     db,
	// }
	// app.DB.AutoMigrate(&models.Person{}, &models.Image{})

	// // Setup routes

	// app.Router.Use(cors.Default())

	// // Define the MEPs endpoint

	// err = app.Router.Run()
	// if err != nil {
	// 	panic(err)
	// }
}
