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
	app.Router.GET("/meps/feed", app.GetMEPsFeedHandler)
	app.Router.GET("/meps/show-current", app.GetMEPsShowCurrentHandler)
	app.Router.GET("/meps/show-incoming", app.GetMEPsShowIncomingHandler)
	app.Router.GET("/meps/show-outgoing", app.GetMEPsShowOutgoingHandler)
}

// MEPs endpoints
// Returns data about the Members of the European Parliament (MEPs)
// ref: https://data.europarl.europa.eu/en/developer-corner/opendata-api

func (app *App) GetMEPsFeedHandler(c *gin.Context) {}

func (app *App) GetMEPsShowCurrentHandler(c *gin.Context) {}

func (app *App) GetMEPsShowIncomingHandler(c *gin.Context) {}

func (app *App) GetMEPsShowOutgoingHandler(c *gin.Context) {}

// Returns data about a specific MEP
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

// Returns data about all MEPs
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
