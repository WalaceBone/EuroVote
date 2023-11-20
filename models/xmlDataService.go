package models

import (
	"encoding/xml"
	"fmt"
	"log"
)

// XML Service Interface that will ingest xml files parse them and persist them to the database
type DataService interface {
	IngestData(rawXMLData string) error
	DumpData() error
}

// XMLDataService is an implementation of the DataService interface.
type XMLDataService struct {
	RollCallVoteResults XMLRollCallVoteResults
}

type XMLRollCallVoteResults struct {
	XMLName xml.Name    `xml:"PV.RollCallVoteResults"`
	Titles  []string    `xml:"RollCallVoteResults.Titles"`
	Result  []XMLResult `xml:"RollCallVote.Result"`
}

type XMLResult struct {
	XMLName     xml.Name        `xml:"RollCallVote"`
	Description string          `xml:"RollCallVote.Description.Text"`
	For         []XMLVoteResult `xml:"Result.For"`
	Against     []XMLVoteResult `xml:"Result.Against"`
	Abstention  []XMLVoteResult `xml:"Result.Abstention"`
	Intentions  []XMLVoteResult `xml:"Intentions"`
}

type XMLVoteResult struct {
	Member []XMLMember `xml:"Member"`
}

type XMLMember struct {
	Name           string `xml:"PoliticalGroup.Member.Name"`
	PoliticalGroup string `xml:"Result.PoliticalGroup.List Identifier"`
}

func (s *XMLDataService) IngestData(rawXMLData []byte) error {

	log.Println("Ingesting data")

	// Define the struct to hold the parsed XML data
	var data XMLRollCallVoteResults

	// log.Println(string(rawXMLData))
	// Parse the XML data
	err := xml.Unmarshal(rawXMLData, &data)
	if err != nil {
		return fmt.Errorf("failed to parse XML data: %v", err)
	}
	log.Println("Data:")
	if len(data.Result) > 0 {
		for i := 0; i < len(data.Result); i++ {
			log.Println("RollCallVote:", i)
			log.Println("Description:", data.Result[i].Description)
			log.Println("For:", len(data.Result[i].For))
			log.Println("Against:", len(data.Result[i].Against))
			log.Println("Abstention:", len(data.Result[i].Abstention))

		}
	} else {
		fmt.Printf("No data found\n")
	}
	// Access the parsed data
	// fmt.Println("Title Text:", data.Titles[0].Text)
	// fmt.Println("Field2:", data.XMLName)

	// Persist the data to the database
	return nil
}

func (s *XMLDataService) DumpData() error {
	// Dump the data from the database
	return nil
}

// Make the endpoints for the service

// func MakeEndpoints(s DataService) Endpoints {
// 	return Endpoints{
// 		IngestDataEndpoint: makeIngestDataEndpoint(s),
// 		DumpDataEndpoint:   makeDumpDataEndpoint(s),
// 	}
// }

// func makeDumpDataEndpoint(s DataService) endpoint.Endpoint {

// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		err := s.DumpData()
// 		return DumpDataResponse{Err: err}, nil
// 	}
// }

// func makeIngestDataEndpoint(s DataService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(IngestDataRequest)
// 		err := s.IngestData(req.RawXMLData)
// 		return IngestDataResponse{Err: err}, nil
// 	}
// }
