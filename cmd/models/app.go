package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func loadConfig() (Config, error) {
	var config Config
	configFile, err := os.Open("config/config.json")
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func (app *App) Initialize() error {
	// Connect to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	app.Config, err = loadConfig()
	if err != nil {
		panic(err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = app.DB.AutoMigrate(&Person{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return err
	}

	// Set up the router
	app.Router = gin.Default()
	app.Router.Use(cors.Default())
	// app.DB.AutoMigrate(&Image{})
	app.SetupRoutes()

	return nil
}

func (app *App) SetupRoutes() {

	// MEPs endpoints
	app.Router.GET("/meps", app.GetMEPsHandler)
	app.Router.GET("/meps/:mepId", app.GetMEPHandler)
	app.Router.GET("/meps/feed", app.GetMEPsFeedHandler)
	app.Router.GET("/meps/show-current", app.GetMEPsShowCurrentHandler)
	app.Router.GET("/meps/show-incoming", app.GetMEPsShowIncomingHandler)
	app.Router.GET("/meps/show-outgoing", app.GetMEPsShowOutgoingHandler)

	// Meetings endpoints
	app.Router.GET("/meetings", app.GetMeetingsHandler)
	// app.Router.GET("/meetings/:event-id", app.GetMeetingHandler)

	//
	// app.Router.GET("/meps/list", app.GetMEPsListHandler)
}

// MEPs endpoints
// Returns data about the Members of the European Parliament (MEPs)
// ref: https://data.europarl.europa.eu/en/developer-corner/opendata-api

// Returns a list of MEPs
func (app *App) GetMEPsListHandler(c *gin.Context) []Person {
	var meps []Person
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
	var mep RDF
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

func InsertOrUpdatePerson(db *gorm.DB, personData *Person) error {
	var existingPerson Person
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

func InsertOrUpdateBatchPerson(db *gorm.DB, personData []Person) error {
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
	var meps RDF
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

	// var existingPerson Person
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

// Meetings endpoints
// Returns data about the meetings of the European Parliament
// ref: https://data.europarl.europa.eu/en/developer-corner/opendata-api

// Returns a list of meetings
// /meetings
// year: The year of the meeting
// offset: The offset of the meetings to return
// limit: The maximum number of meetings to return
func (app *App) GetMeetingsHandler(c *gin.Context) {
	apiUrl := app.Config.EuroparlAPIURL + "/meetings" // Use the config URL

	// Make an HTTP request
	// year := c.Query("2020")
	// offset := c.Query("10")
	// limit := c.Query("10")
	newreq, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	q := newreq.URL.Query()
	q.Add("year", "2022")
	q.Add("offset", "0")
	q.Add("limit", "10")
	newreq.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(newreq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	defer resp.Body.Close()

	// Read and print the response body

	log.Println(resp.Request.URL)
	body, err := io.ReadAll(resp.Body)
	log.Println("Response body is:")
	log.Println(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(http.StatusOK, "application/ld+json", body)
}

func (app *App) LoadAndSavePersonData(c *gin.Context) error {
	var aboutPerson RDF

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
