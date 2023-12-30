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

	// app.Router.GET("/meps/list", app.GetMEPsListHandler)
}

// MEPs endpoints
// Returns data about the Members of the European Parliament (MEPs)
// ref: https://data.europarl.europa.eu/en/developer-corner/opendata-api

// Returns a list of MEPs
func (app *App) GetMEPsListHandler(c *gin.Context) []models.Person {
	var meps []models.Person
	result := app.DB.Find(&meps)

	fmt.Println(result.RowsAffected)
	fmt.Println(len(meps))
	if result.Error != nil {
		fmt.Println("Error is:")
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return nil
	}

	c.JSON(http.StatusOK, meps)
	return meps
}

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

func InsertOrUpdatePerson(db *gorm.DB, personData *models.Person) error {
	var existingPerson models.Person
	// Try to find the record
	log.Println("Identifier is:")
	log.Println(personData.Identifier)
	log.Println("ID is:")
	log.Println(personData.ID)
	result := db.Where("ID = ?", personData.ID).First(&existingPerson)
	log.Println("Result is:")
	log.Println(result)
	if result.Error == gorm.ErrRecordNotFound {
		// Record not found, create a new one
		if createResult := db.Create(personData); createResult.Error != nil {
			return createResult.Error // Handle create error
		}
		log.Println("New person is:")
		log.Println(personData)
	} else {
		// Record found, update it
		existingPerson.About = personData.About
		existingPerson.SortLabel = personData.SortLabel
		existingPerson.Label = personData.Label
		existingPerson.Identifier = personData.Identifier
		existingPerson.GivenName = personData.GivenName
		existingPerson.FamilyName = personData.FamilyName
		existingPerson.Memberships = personData.Memberships
		existingPerson.CreatedAt = personData.CreatedAt
		existingPerson.UpdatedAt = personData.UpdatedAt
		existingPerson.Birthday = personData.Birthday
		existingPerson.Citizenship = personData.Citizenship
		existingPerson.DeathDate = personData.DeathDate
		existingPerson.HonorificPrefix = personData.HonorificPrefix
		existingPerson.Gender = personData.Gender
		existingPerson.PlaceOfBirth = personData.PlaceOfBirth
		existingPerson.UpperFamilyName = personData.UpperFamilyName
		existingPerson.UpperGivenName = personData.UpperGivenName
		existingPerson.Notation = personData.Notation

		// Update other fields as necessary
		if updateResult := db.Save(&existingPerson); updateResult.Error != nil {
			return updateResult.Error // Handle update error
		}
		log.Println("Existing person is:")
		log.Println(existingPerson)
	}

	return nil
}

func InsertOrUpdateBatchPerson(db *gorm.DB, personData []models.Person) error {
	return db.CreateInBatches(personData, 100).Error
}

// Returns data about all MEPs
// /meps
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

	err = InsertOrUpdateBatchPerson(app.DB, meps.Person)
	if err != nil {
		return
	}

	// for _, person := range meps.Person {
	// // Add id and creationDate fields to person
	// // person.ID = generateID() // Replace generateID() with your own logic to generate unique IDs
	// person.CreatedAt = time.Now()

	// var existingPerson models.Person
	// result := app.DB.Where("identifier = ?", person.Identifier).First(&existingPerson)

	// log.Println(existingPerson.ID)
	// if result.Error == gorm.ErrRecordNotFound {

	// 	err = InsertOrUpdateBatchPerson(app.DB, &person)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	// result := app.DB.Create(&person)
	// if result.Error != nil {
	// 	if result.Error == gorm.ErrRecordNotFound {
	// 		// Person does not exist, create it

	// 		// result = app.DB.Create(&person)

	// 		if result.Error != nil {
	// 			fmt.Println("Error is:")
	// 			fmt.Println(result.Error)
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	// 			return
	// 		}
	// 	} else {
	// 		// Error occurred while querying the database
	// 		fmt.Println("Error is:")
	// 		fmt.Println(result.Error)

	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	// 		return
	// 	}
	// } else {
	// 	// Person already exists
	// 	fmt.Println("Person already exists")
	// }

	// if result.Error != nil {
	// 	fmt.Println("Error is:")
	// 	fmt.Println(result.Error)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	// 	return
	// }
	// }
	app.LoadAndSavePersonData(c)

	app.GetMEPsListHandler(c)
	// c.JSON(http.StatusOK, meps.Person)
}

func (app *App) LoadAndSavePersonData(c *gin.Context) error {
	var aboutPerson models.RDF

	listPerson := app.GetMEPsListHandler(c)
	for _, person := range listPerson {

		res, err := http.Get(person.About)
		log.Println("Person about is:")
		log.Println(person)

		if err != nil {
			log.Println("Error is:")
			log.Println(err)
			return err
		}
		defer res.Body.Close()

		// Read and parse the response

		if err := xml.NewDecoder(res.Body).Decode(&aboutPerson); err != nil {
			log.Println("Error is:")
			log.Println(err)
			return err
		}
		for _, about := range aboutPerson.Person {
			about.UpdatedAt = time.Now()
			InsertOrUpdatePerson(app.DB, &about)

		}

	}

	return nil
}
