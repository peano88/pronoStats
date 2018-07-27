package dataLayer

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Prono struct {
	ID             bson.ObjectId
	HomeTeam       string
	AwayTeam       string
	HomeScore      int
	AwayScore      int
	PronoHomeScore int
	PronoAwayScore int
}

type DataBridge struct {
	Coll *mgo.Collection
}

const (
	PRONO_COLLECTION = "PRONOS"
)

func (db *DataBridge) AddProno(pr Prono) {
}

func (db *DataBridge) FindPronoById(id string) (Prono, error) {
	var prono Prono

	return prono, nil
}
