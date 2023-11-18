package model

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// VoteResult represents the vote result data model.
type VoteResult struct {
	gorm.Model
	// Add your fields here based on your data structure
	Identifier string
	DlvID      string
	Date       string
	// ... other fields
}

type App struct {
	DB     *gorm.DB
	Client *http.Client
	Router *gin.Engine
	// Add other fields as needed
}

// Run starts the application.
func (a *App) Run() {
	a.Router.Run()
}

// Stop stops the application.
func (a *App) Stop() {
	a.DB.Close()
}

type PoliticalGroup struct {
	gorm.Model
	GroupID    uint `gorm:"primaryKey"`
	Identifier string
	Members    []Member
}

type Member struct {
	gorm.Model
	MemberID       uint `gorm:"primaryKey"`
	MepID          string
	PersID         string
	GroupID        uint
	PoliticalGroup PoliticalGroup `gorm:"foreignKey:GroupID"`
	RollCallVotes  []RollCallVote
}

type RollCallVote struct {
	gorm.Model
	VoteID     uint `gorm:"primaryKey"`
	Identifier string
	DlvID      string
	Date       time.Time
	MemberID   uint
	Member     Member `gorm:"foreignKey:MemberID"`
	VoteResult string
}

func (app *App) CreateGroup(identifier string) *PoliticalGroup {
	group := &PoliticalGroup{Identifier: identifier}
	app.DB.Create(group)
	return group
}

func (app *App) GetGroup(id uint) *PoliticalGroup {
	var group PoliticalGroup
	app.DB.First(&group, id)
	return &group
}

func (app *App) UpdateGroup(id uint, identifier string) *PoliticalGroup {
	group := app.GetGroup(id)
	group.Identifier = identifier
	app.DB.Save(group)
	return group
}

func (app *App) DeleteGroup(id uint) {
	group := app.GetGroup(id)
	app.DB.Delete(group)
}

func (app *App) CreateMember(mepID string, persID string, groupID uint) *Member {
	member := &Member{MepID: mepID, PersID: persID, GroupID: groupID}
	app.DB.Create(member)
	return member
}

func (app *App) GetMember(id uint) *Member {
	var member Member
	app.DB.First(&member, id)
	return &member
}

func (app *App) UpdateMember(id uint, mepID string, persID string) *Member {
	member := app.GetMember(id)
	member.MepID = mepID
	member.PersID = persID
	app.DB.Save(member)
	return member
}

func (app *App) DeleteMember(id uint) {
	member := app.GetMember(id)
	app.DB.Delete(member)
}

func (app *App) CreateVote(identifier string, dlvID string, date time.Time, memberID uint, voteResult string) *RollCallVote {
	vote := &RollCallVote{Identifier: identifier, DlvID: dlvID, Date: date, MemberID: memberID, VoteResult: voteResult}
	app.DB.Create(vote)
	return vote
}

func (app *App) GetVote(id uint) *RollCallVote {
	var vote RollCallVote
	app.DB.First(&vote, id)
	return &vote
}

func (app *App) UpdateVote(id uint, identifier string, voteResult string) *RollCallVote {
	vote := app.GetVote(id)
	vote.Identifier = identifier
	vote.VoteResult = voteResult
	app.DB.Save(vote)
	return vote
}

func (app *App) DeleteVote(id uint) {
	vote := app.GetVote(id)
	app.DB.Delete(vote)
}

// Similar functions can be created for Member and RollCallVote
