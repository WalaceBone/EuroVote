package main

import (
	"encoding/xml"
	"fmt"

	"github.com/WalaceBone/EuroVote/models"
)

// XMLDataServiceImpl is an implementation of the DataService interface.
type XMLDataServiceImpl struct {
}

func (s *XMLDataServiceImpl) IngestData(rawXMLData string) error {
	// Parse the raw XML data
	var voteResults models.RollCallVoteResults
	if err := xml.Unmarshal([]byte(rawXMLData), &voteResults); err != nil {
		return err
	}

	// Process the parsed data (you can add your logic here)
	fmt.Println("Processed XML data:", voteResults)

	// Your additional processing logic goes here

	return nil
}

func main() {
	// Example usage

	app, err := initApp()

	if err != nil {
		fmt.Println("Error:", err)
	}

	app.Router.StaticFile("/", "./public/index.html")

	app.Run()
	defer app.Stop()

}
