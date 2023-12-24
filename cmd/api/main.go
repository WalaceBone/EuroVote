package main

import (
	"EuroVote/cmd/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func loadConfig() (models.Config, error) {
	var config models.Config
	configFile, err := os.Open("config/config.json")
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main() {

	fmt.Println("Starting the application...")
	fmt.Println(os.Getenv("PORT"))

	Config, err := loadConfig()

	if err != nil {
		panic(err)
	}

	app := App{
		Router: gin.Default(),
		Config: Config,
	}

	// Setup routes

	app.Router.Use(cors.Default())
	app.SetupRoutes()

	// Define the MEPs endpoint

	err = app.Router.Run()
	if err != nil {
		panic(err)
	}
}
