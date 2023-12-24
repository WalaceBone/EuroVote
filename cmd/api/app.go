package main

import (
	"EuroVote/cmd/models"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	Config models.Config
	DB     *gorm.DB
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

// Returns an Atom feed of MEPs
func (app *App) GetMEPsFeedHandler(c *gin.Context) {

	apiUrl := app.Config.EuroparlAPIURL + "/meps/feed" // Use the config URL

	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	defer resp.Body.Close()

	// Read and process the response body
	// This step will depend on how the data is returned by the European Parliament's API
	// You might need to transform it into the Atom feed format
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// If additional processing is required to convert the response to Atom XML format, do it here

	// Return the feed
	c.Data(http.StatusOK, "application/atom+xml", body)
}

// Returns data about the current MEPs
func (app *App) GetMEPsShowCurrentHandler(c *gin.Context) {
	// apiUrl := app.Config.EuroparlAPIURL + "/meps/show-current" // Use the config URL

	gender := c.Query("gender")
	politicalGroup := c.Query("political-group")
	country := c.Query("country-of-representation")
	format := c.Query("format-all")
	offset := c.Query("offset")
	limit := c.Query("limit")

	apiUrl := fmt.Sprintf(app.Config.EuroparlAPIURL+"/meps/show-current?date=%s&gender=%s&group=%s&country=%s&format=%s&offset=%s&limit=%s",
		time.Now().Format("2006-01-02"), gender, politicalGroup, country, format, offset, limit)

	// Make an HTTP request
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Forward the API's response to the client
	// You can add additional processing here if needed
	c.Data(resp.StatusCode, "application/ld+json", body)

}

// Returns data about the incoming MEPs
func (app *App) GetMEPsShowIncomingHandler(c *gin.Context) {
	// Parse query parameters
	gender := c.Query("gender")
	politicalGroup := c.Query("political-group")
	country := c.Query("country-of-representation")
	format := c.Query("format-all")
	offset := c.Query("offset")
	limit := c.Query("limit")

	// Construct the API request to the European Parliament's API
	// The URL structure and parameters might need to be adjusted based on the actual API documentation
	apiUrl := fmt.Sprintf(app.Config.EuroparlAPIURL+"/meps/show-incoming?gender=%s&group=%s&country=%s&format=%s&offset=%s&limit=%s",
		gender, politicalGroup, country, format, offset, limit)

	// Fetch the data from the European Parliament's API
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Forward the API's response to the client
	c.Data(resp.StatusCode, "application/ld+json", body)
}

// Returns data about the outgoing MEPs
func (app *App) GetMEPsShowOutgoingHandler(c *gin.Context) {
	// Parse query parameters
	gender := c.Query("gender")
	politicalGroup := c.Query("political-group")
	country := c.Query("country-of-representation")
	format := c.Query("format-all")
	offset := c.Query("offset")
	limit := c.Query("limit")

	// Construct the API request to the European Parliament's API
	apiUrl := fmt.Sprintf(app.Config.EuroparlAPIURL+"/meps/show-outgoing?gender=%s&group=%s&country=%s&format=%s&offset=%s&limit=%s",
		gender, politicalGroup, country, format, offset, limit)

	// Fetch the data from the European Parliament's API
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	defer resp.Body.Close()

	// Read and process the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Forward the API's response to the client
	c.Data(resp.StatusCode, "application/ld+json", body)
}

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

	result := app.DB.Create(&mep.Person[0])
	if result.Error != nil {
		fmt.Println("Error is:")
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
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

	for _, person := range meps.Person {
		result := app.DB.Create(&person)
		if result.Error != nil {
			fmt.Println("Error is:")
			fmt.Println(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, meps.Person)
}
