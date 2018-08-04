package dataLayer

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Prono struct {
	HomeTeam       string `bson:"home_team" json:"home_team"`
	AwayTeam       string `bson:"away_team" json:"away_team"`
	HomeScore      int    `bson:"home_score" json:"home_score"`
	AwayScore      int    `bson:"away_score" json:"away_score"`
	PronoHomeScore int    `bson:"prono_home_team" json:"prono_home_team"`
	PronoAwayScore int    `bson:"prono_away_score" json:"prono_away_score"`
}

type TournamentPronos struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Sport      string        `bson:"sport" json:"sport"`
	Tournament string        `bson:"tournament" json:"tournament"`
	Pronos     []Prono       `bson:"pronos" json:"pronos"`
}

type DataBridge struct {
	Coll *mgo.Collection
}

const (
	PRONO_COLLECTION = "PRONOS"
	DB_NAME          = "PRONO_STATS"
)

func (db *DataBridge) AddTournamentPronos(tp TournamentPronos) (string, error) {
	if tp.ID == "" {
		tp.ID = bson.NewObjectId()
	}

	if err := db.Coll.Insert(&tp); err != nil {
		return "", err
	}

	return tp.ID.Hex(), nil
}

func (db *DataBridge) FindTournamentPronosById(id string) (TournamentPronos, error) {
	var tPronos TournamentPronos

	if !bson.IsObjectIdHex(id) {
		return tPronos, errors.New("invalid id")
	}
	err := db.Coll.FindId(bson.ObjectIdHex(id)).One(&tPronos)

	return tPronos, err
}

func (db *DataBridge) AddProno(idTourPronos string, pr Prono) error {
	if !bson.IsObjectIdHex(idTourPronos) {
		return errors.New("invalid id")
	}
	return db.Coll.UpdateId(idTourPronos, bson.M{"$push": &pr})

}

//Disabled
/*
func (db *DataBridge) FindPronoById(id string) (Prono, error) {
	var prono Prono
	if !bson.IsObjectIdHex(id) {
		return prono, errors.New("invalid id")
	}
	err := db.Coll.FindId(bson.ObjectIdHex(id)).One(&prono)

	return prono, err
}
*/
