package dataLayer

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Team struct {
	Name       string   `bson:"name" json:"name"`
	Categories []string `bson:"categories" json:"categories"`
}

type Match struct {
	HomeTeam  string `bson:"home_team" json:"home_team"`
	AwayTeam  string `bson:"away_team" json:"away_team"`
	HomeScore int    `bson:"home_score" json:"home_score"`
	AwayScore int    `bson:"away_score" json:"away_score"`
}

type Round struct {
	Number  int     `bson:"number" json:"number"`
	Matches []Match `bson:"matches" json:"matches"`
}

type Tournament struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Sport  string        `bson:"sport" json:"sport"`
	Teams  []Team        `bson:"teams" json:"teams"`
	Rounds []Round       `bson:"rounds" json:"rounds"`
}

func (db *DataBridge) AddTournament(t Tournament) (string, error) {
	if t.ID == "" {
		t.ID = bson.NewObjectId()
	}

	if err := db.TournColl.Insert(&t); err != nil {
		return "", err
	}

	return t.ID.Hex(), nil
}

func (db *DataBridge) FindTournamentById(id string) (Tournament, error) {
	var t Tournament

	if !bson.IsObjectIdHex(id) {
		return t, errors.New("invalid id")
	}
	err := db.TournColl.FindId(bson.ObjectIdHex(id)).One(&t)

	return t, err
}
