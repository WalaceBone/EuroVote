package models

import (
	"encoding/xml"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	XMLName    xml.Name `xml:"Person" gorm:"-"`
	About      string   `xml:"about,attr" gorm:"type:text"`
	SortLabel  string   `xml:"sortLabel" gorm:"type:text"`
	Label      string   `xml:"label" gorm:"type:text"`
	Identifier string   `xml:"identifier" gorm:"type:text"`
	GivenName  string   `xml:"givenName" gorm:"type:text"`
	FamilyName string   `xml:"familyName" gorm:"type:text"`
}

type RDF struct {
	XMLName xml.Name `xml:"RDF"`
	Person  []Person `xml:"Person"`
}

type Config struct {
	EuroparlAPIURL string `json:"europarl_api_url"`
}
