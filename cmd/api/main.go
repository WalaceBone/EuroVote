package main

import (
	"EuroVote/cmd/models"
	"encoding/json"
	"os"

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

	Config, err := loadConfig()

	if err != nil {
		panic(err)
	}

	app := App{
		Router: gin.Default(),
		Config: Config,
	}

	// Setup routes

	app.SetupRoutes()

	// Define the MEPs endpoint

	app.Router.Run()

}
