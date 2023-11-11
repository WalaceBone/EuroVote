package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PoliticalGroup struct {
	gorm.Model
	GroupID    uint   `gorm:"primary_key"`
	Identifier string `gorm:"type:varchar(100)"`
	Members    []Member
}

type Member struct {
	gorm.Model
	MemberID  uint   `gorm:"primary_key"`
	MepID     string `gorm:"type:varchar(100)"`
	PersID    string `gorm:"type:varchar(100)"`
	GroupID   uint
	RollCalls []RollCallVote
}

type RollCallVote struct {
	gorm.Model
	VoteID     uint      `gorm:"primary_key"`
	Identifier string    `gorm:"type:varchar(100)"`
	DlvID      string    `gorm:"type:varchar(100)"`
	Date       time.Time `gorm:"type:timestamp"`
	MemberID   uint
	VoteResult string `gorm:"type:varchar(100)"`
}

type DB struct {
	*gorm.DB
}

var db *DB

func start() {
	var err error
	dsn := "host=127.0.0.1 user=postgres password=mysecretpassword dbname=postgres port=5432"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	db = &DB{dbConn}
}

func main() {

	start()
	// Get the absolute path to the XML file
	xmlPath := filepath.Join(".", "public", "xml", "vote.xml")
	absPath, err := filepath.Abs(xmlPath)
	if err != nil {
		log.Fatal(err)
	}

	// Read the contents of the file into a byte slice
	xmlData, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of the top-level struct
	print(xmlData)
	// Do something with the result...
}
