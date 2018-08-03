package dataLayer

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Prono struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	HomeTeam       string        `bson:"home_team" json:="home_team"`
	AwayTeam       string        `bson:"away_team" json:="away_team"`
	HomeScore      int           `bson:"home_score" json:="home_score"`
	AwayScore      int           `bson:"away_score" json:="away_score"`
	PronoHomeScore int           `bson:"prono_home_team" json:="prono_home_team"`
	PronoAwayScore int           `bson:"prono_away_score" json:="prono_away_score"`
}

type DataBridge struct {
	Coll *mgo.Collection
}

const (
	PRONO_COLLECTION = "PRONOS"
	DB_NAME          = "PRONO_STATS"
)

func (db *DataBridge) AddProno(pr Prono) (string, error) {
	if pr.ID == "" {
		pr.ID = bson.NewObjectId()
	}
	if err := db.Coll.Insert(&pr); err != nil {
		return "", err
	}

	return pr.ID.Hex(), nil
}

func (db *DataBridge) FindPronoById(id string) (Prono, error) {
	var prono Prono
	if !bson.IsObjectIdHex(id) {
		return prono, errors.New("invalid id")
	}
	err := db.Coll.FindId(bson.ObjectIdHex(id)).One(&prono)

	return prono, err
}
