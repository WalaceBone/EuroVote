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
	Titles  []Title     `xml:"RollCallVoteResults.Titles>RollCallVoteResults.Title.Text"`
	Result  []XMLResult `xml:"RollCallVote.Result"`
}

type Title struct {
	Text     string `xml:",chardata"`
	Language string `xml:"Language,attr"`
}

type XMLResult struct {
	XMLName     xml.Name `xml:"RollCallVote.Result"`
	Description string   `xml:"Description.Text"`
	For         XMLVote  `xml:"Result.For"`
	Against     XMLVote  `xml:"Result.Against"`
	Abstention  XMLVote  `xml:"Result.Abstention"`
	Intentions  string   `xml:"Intentions Number,attr"`
}

type XMLVote struct {
	Count          int                 `xml:"Number,attr"`
	PoliticalGroup []XMLPoliticalGroup `xml:"Result.PoliticalGroup.List"`
}

type XMLPoliticalGroup struct {
	// XMLName   xml.Name    	`xml:"PoliticalGroup.List"`
	Identifier string      `xml:"Identifier,attr"`
	Members    []XMLMember `xml:"PoliticalGroup.Member.Name"`
}

type XMLMember struct {
	Name   string `xml:",chardata"`
	MepId  int    `xml:"MepId,attr"`
	PersId int    `xml:"PersId,attr"`
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
			for j := 0; j < len(data.Titles); j++ {
				log.Println("Title:", data.Titles[j].Text)
				log.Println("Language:", data.Titles[j].Language)
			}
			log.Println("Description:", data.Result[i].Description)
			log.Println("For:", data.Result[i].For)
			log.Println("Count:", data.Result[i].For.Count)
			log.Println("PoliticalGroup:", data.Result[i].For.PoliticalGroup)
			for k := 0; k < len(data.Result[i].For.PoliticalGroup); k++ {
				log.Println("PoliticalGroup:", data.Result[i].For.PoliticalGroup[k].Identifier)
				for l := 0; l < len(data.Result[i].For.PoliticalGroup[k].Members); l++ {
					log.Println("Member:", data.Result[i].For.PoliticalGroup[k].Members[l].Name)
					log.Println("MepId:", data.Result[i].For.PoliticalGroup[k].Members[l].MepId)
					log.Println("PersId:", data.Result[i].For.PoliticalGroup[k].Members[l].PersId)
				}
			}
			log.Println("Against:", data.Result[i].Against)
			log.Println("Count:", data.Result[i].Against.Count)
			for m := 0; m < len(data.Result[i].Against.PoliticalGroup); m++ {
				log.Println("PoliticalGroup:", data.Result[i].Against.PoliticalGroup[m].Identifier)
				for n := 0; n < len(data.Result[i].Against.PoliticalGroup[m].Members); n++ {
					log.Println("Member:", data.Result[i].Against.PoliticalGroup[m].Members[n].Name)
					log.Println("MepId:", data.Result[i].Against.PoliticalGroup[m].Members[n].MepId)
					log.Println("PersId:", data.Result[i].Against.PoliticalGroup[m].Members[n].PersId)
				}
			}

			log.Println("Abstention:", data.Result[i].Abstention)
			log.Println("Count:", data.Result[i].Abstention.Count)
			for k := 0; k < len(data.Result[i].Abstention.PoliticalGroup); k++ {
				log.Println("PoliticalGroup:", data.Result[i].Abstention.PoliticalGroup[k].Identifier)
				for l := 0; l < len(data.Result[i].Abstention.PoliticalGroup[k].Members); l++ {
					log.Println("Member:", data.Result[i].Abstention.PoliticalGroup[k].Members[l].Name)
					log.Println("MepId:", data.Result[i].Abstention.PoliticalGroup[k].Members[l].MepId)
					log.Println("PersId:", data.Result[i].Abstention.PoliticalGroup[k].Members[l].PersId)
				}
			}
		}
	} else {
		log.Println("No data")
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
