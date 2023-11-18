package main

import (
	model "github.com/WalaceBone/EuroVote/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=mysecretpassword")
	if err != nil {
		return nil, err
	}

	// AutoMigrate will create tables based on your GORM model
	db.AutoMigrate(&model.VoteResult{})
	return db, nil
}

func initRouter() *gin.Engine {
	r := gin.Default()
	// Add your routes here

	r.GET("/hello", func(c *gin.Context) {

		c.JSON(200, "Hello World!")
	})

	return r
}

func initApp() (*model.App, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	app := &model.App{DB: db}

	app.DB, err = initDB()
	if err != nil {
		return nil, err
	}

	app.Router = initRouter()
	return app, nil
}
