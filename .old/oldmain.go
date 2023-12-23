package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName    xml.Name `xml:"Person"`
	About      string   `xml:"about,attr"`
	SortLabel  string   `xml:"sortLabel"`
	Label      string   `xml:"label"`
	Identifier string   `xml:"identifier"`
	GivenName  string   `xml:"givenName"`
	FamilyName string   `xml:"familyName"`
}

type RDF struct {
	XMLName xml.Name `xml:"RDF"`
	Person  []Person `xml:"Person"`
}

func tmp() {
	// Example usage

	router := gin.Default()

	// Define the MEPs endpoint
	router.GET("/meps", func(c *gin.Context) {
		// Endpoint of the European Parliament API for MEPs
		apiUrl := "https://data.europarl.europa.eu/api/v1/meps"

		// Send GET request to the European Parliament API
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

		var rdf RDF
		err = xml.Unmarshal([]byte(body), &rdf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
			return
		}

		// fmt.Println(rdf)

		// Sort the rdf.Person slice by Identifier
		sort.Slice(rdf.Person, func(i, j int) bool {
			return rdf.Person[i].Identifier < rdf.Person[j].Identifier
		})

		for _, person := range rdf.Person {
			fmt.Printf("Person struct: %+v\n", person)
		}

		if len(rdf.Person) > 0 {
			c.JSON(http.StatusOK, len(rdf.Person))
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "MEP not found"})
		}
	})

	// Endpoint for fetching data about a specific MEP
	router.GET("/meps/:mepId", func(c *gin.Context) {
		mepId := c.Param("mepId")

		// Construct the URL to fetch data from the European Parliament API
		apiUrl := fmt.Sprintf("https://data.europarl.europa.eu/api/v1/meps/%s", mepId)

		// Fetch the data
		resp, err := http.Get(apiUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		var rdf RDF
		err = xml.Unmarshal(body, &rdf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
			return
		}

		for _, person := range rdf.Person {
			fmt.Printf("Person struct: %+v\n", person)
		}

		// Assuming that the API returns only one Person per ID
		if len(rdf.Person) > 0 {
			c.JSON(http.StatusOK, rdf.Person)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "MEP not found"})
		}
	})

	// Define the EPBodies endpoint
	router.GET("/epbodies", func(c *gin.Context) {
		// Mock data
		epbodies := []map[string]interface{}{
			{"id": 1, "name": "EPBody 1"},
			{"id": 2, "name": "EPBody 2"},
		}
		c.JSON(http.StatusOK, epbodies)
	})

	// Define the EPDocuments endpoint
	router.GET("/epdocuments", func(c *gin.Context) {
		// Mock data
		epdocuments := []map[string]interface{}{
			{"id": 1, "title": "Document 1"},
			{"id": 2, "title": "Document 2"},
		}
		c.JSON(http.StatusOK, epdocuments)
	})

	// Define the EPEvents endpoint
	router.GET("/epevents", func(c *gin.Context) {
		// Mock data
		epevents := []map[string]interface{}{
			{"id": 1, "name": "Event 1"},
			{"id": 2, "name": "Event 2"},
		}
		c.JSON(http.StatusOK, epevents)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.Run()

}
