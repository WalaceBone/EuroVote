package models

import "encoding/xml"

type RollCallVoteResults struct {
	XMLName       xml.Name          `xml:"RollCallVoteResults.Titles"`
	Titles        []Title           `xml:"Title.Text"`
	RollCallVote  []XMLRollCallVote `xml:"RollCallVote.Result"`
	XMLVoteResult []XMLVoteResult   `xml:"RollCallVote.Result.Result"`
}

type Title struct {
	Language string `xml:"Language,attr"`
	Text     string `xml:",chardata"`
}

type XMLRollCallVote struct {
	XMLName     xml.Name    `xml:"RollCallVote.Result"`
	Identifier  string      `xml:"Identifier,attr"`
	DlvId       string      `xml:"DlvId,attr"`
	Date        string      `xml:"Date,attr"`
	Description Description `xml:"Description.Text"`
	Result      Result      `xml:"Result"`
	Intentions  Intentions  `xml:"Intentions"`
}

type Description struct {
	Text string `xml:",chardata"`
}

type Result struct {
	For        VoteResult `xml:"For"`
	Against    VoteResult `xml:"Against"`
	Abstention VoteResult `xml:"Abstention"`
}

type XMLVoteResult struct {
	Number             string               `xml:"Number,attr"`
	PoliticalGroupList []PoliticalGroupList `xml:"PoliticalGroup.List"`
}

type PoliticalGroupList struct {
	Identifier string `xml:"Identifier,attr"`
	MemberName string `xml:"Member.Name"`
}

type Intentions struct {
	Code       string      `xml:"Code,attr"`
	For        []XMLMember `xml:"Result.For>Member.Name"`
	Against    []XMLMember `xml:"Result.Against>Member.Name"`
	Abstention []XMLMember `xml:"Result.Abstention>Member.Name"`
}
type XMLMember struct {
	MepId  string `xml:"MepId,attr"`
	PersId string `xml:"PersId,attr"`
	Name   string `xml:",chardata"`
}
