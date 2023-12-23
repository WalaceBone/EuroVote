package main

type Mep struct {
	ID                         string `json:"id"`
	Type                       string `json:"type"`
	SortLabel                  string `json:"sortLabel"`
	Identifier                 string `json:"identifier"`
	Label                      string `json:"label"`
	GivenName                  string `json:"givenName"`
	FamilyName                 string `json:"familyName"`
	OfficialGivenName          string `json:"officialGivenName"`
	OfficialFamilyName         string `json:"officialFamilyName"`
	APICountryOfRepresentation string `json:"api:country-of-representation"`
	APIPoliticalGroup          string `json:"api:political-group"`
}

type MEPs struct {
	Data []Mep `json:"data"`
}

type EPBody struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	Label      string `json:"label"`
	LinkedTo   string `json:"linkedTo"`
}

type EPBodies struct {
	Data []EPBody `json:"data"`
}

type EPDocument struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	Label      string `json:"label"`
	WorkType   string `json:"work_type"`
}

type EPDocuments struct {
	Data []EPDocument `json:"data"`
}

type EPEvent struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ActivityID      string `json:"activity_id"`
	ActivityLabel   string `json:"activity_label"`
	HadActivityType string `json:"had_activity_type"`
}

type EPEvents struct {
	Data []EPEvent `json:"data"`
}
