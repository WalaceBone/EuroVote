package main

import (
	"EuroVote/cmd/models"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config models.Config
}

func (app *App) SetupRoutes() {
	app.Router.GET("/meps", app.GetMEPsHandler)
	app.Router.GET("/meps/:mepId", app.GetMEPHandler)
}

func (app *App) GetMEPHandler(c *gin.Context) {
	mepId := c.Param("mepId")

	apiUrl := app.Config.EuroparlAPIURL + "/meps/" + mepId // Use the config URL

	// Make an HTTP request
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read and parse the response
	var mep models.RDF
	if err := xml.NewDecoder(resp.Body).Decode(&mep); err != nil {
		log.Println("Error is:")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mep.Person[0])
}

func (app *App) GetMEPsHandler(c *gin.Context) {
	apiUrl := app.Config.EuroparlAPIURL + "/meps" // Use the config URL

	// Make an HTTP request
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read and parse the response
	var meps models.RDF
	if err := xml.NewDecoder(resp.Body).Decode(&meps); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meps.Person)
}
