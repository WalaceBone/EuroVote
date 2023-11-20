package models

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App struct {
	DB     *gorm.DB
	Client *http.Client
	Router *gin.Engine
	// Add other fields as needed
}

func (a *App) SetupRoutes() {
	a.Router.GET("/group/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		group := a.GetGroup(uint(id))
		c.JSON(200, group)
	})

	a.Router.GET("/member/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		member := a.GetMember(uint(id))
		c.JSON(200, member)
	})

	a.Router.GET("/vote/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		vote := a.GetVote(uint(id))
		c.JSON(200, vote)
	})

	// Route to handle the dump request that will respond with the data from the database
	a.Router.GET("/dump", func(c *gin.Context) {

		// Dump the data from the database
		// Send the data back to the client
	})

	a.Router.POST("/xml", func(c *gin.Context) {
		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		openFile, err := file.Open()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer openFile.Close()

		fc, err := io.ReadAll(openFile)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// fmt.Println("xml :", string(fc))

		// fmt.Print("HELLO\n")
		//Basic success response
		// Send the file to the XMLService
		xmlService := &XMLDataService{}
		err = xmlService.IngestData(fc)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})

	})

}

// Run starts the application.
func (a *App) Run() {
	a.Router.Run()
}

// Stop stops the application.
func (a *App) Stop() {
	a.DB.Close()
}
