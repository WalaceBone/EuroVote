package models

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActivityType string

const (
	EPMeetingPart        ActivityType = "EP_MEETING_PART"
	EPPlenarySitting     ActivityType = "EP_PLENARY_SITTING"
	EPPlenaryPartSession ActivityType = "EP_PLENARY_PART_SESSION"
)

type Person struct {
	gorm.Model
	XMLName     xml.Name `xml:"Person" gorm:"-"`
	About       string   `xml:"about,attr" gorm:"type:text"`
	SortLabel   string   `xml:"sortLabel" gorm:"type:text"`
	Label       string   `xml:"label" gorm:"type:text"`
	Identifier  string   `xml:"identifier" gorm:"unique,type:text"`
	GivenName   string   `xml:"givenName" gorm:"type:text"`
	FamilyName  string   `xml:"familyName" gorm:"type:text"`
	Memberships []string `xml:"hasMembership" gorm:"-"`
	Gender      string   `xml:"hasGender" gorm:"type:text"`
	// Image           Image    `xml:"Image"`
	Notation        string `xml:"notation" gorm:"type:text"`
	PlaceOfBirth    string `xml:"placeOfBirth" gorm:"type:text"`
	UpperFamilyName string `xml:"upperFamilyName" gorm:"type:text"`
	DeathDate       string `xml:"deathDate" gorm:"type:text"`
	UpperGivenName  string `xml:"upperGivenName" gorm:"type:text"`
	Citizenship     string `xml:"citizenship" gorm:"type:text"`
	HonorificPrefix string `xml:"hasHonorificPrefix" gorm:"type:text"`
	Birthday        string `xml:"bday" gorm:"type:text"`
}

type RDF struct {
	XMLName xml.Name `xml:"RDF"`
	Person  []Person `xml:"Person"`
}

type Config struct {
	EuroparlAPIURL string `json:"europarl_api_url"`
	APIURL         string `json:"api_url"`
}

type Image struct {
	gorm.Model
	XMLName xml.Name `xml:"Image" gorm:"-"`
	About   string   `xml:"about,attr" gorm:"type:text"`
}

type App struct {
	Router *gin.Engine
	Config Config
	DB     *gorm.DB
}

type EPEvents struct {
	Data []EPEventData `json:"data"`
}

type EPEventData struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ActivityID      string `json:"activity_id"`
	ActivityLabel   string `json:"activity_label"`
	HadActivityType string `json:"had_activity_type"`
}
