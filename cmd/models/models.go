package models

import (
	"encoding/xml"
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

type Config struct {
	EuroparlAPIURL string `json:"europarl_api_url"`
}
