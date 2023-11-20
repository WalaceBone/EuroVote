package models

import (
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

	// Route receiving a XML file to send to a XMLService
	// that will parse it and persist it to the database
	a.Router.POST("/xml", func(c *gin.Context) {
		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the file locally
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Basic success response
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		// Send the file to the XMLService
		// xmlService := &XMLDataServiceImpl{}
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
