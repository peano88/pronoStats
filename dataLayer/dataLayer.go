package dataLayer

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type DataBridge struct {
	PronoColl *mgo.Collection
	TournColl *mgo.Collection
}

const (
	PRONO_COLLECTION = "PRONOS"
	TOURN_COLLECTION = "TOURN"
	DB_NAME          = "PRONO_STATS"
)

func (db *DataBridge) Init(session *mgo.Session) {

	dataBase := session.DB(DB_NAME)
	if dataBase == nil {
		log.Fatal("Fatal error in instantiating the DB")
	}

	db.PronoColl = dataBase.C(PRONO_COLLECTION)
	db.TournColl = dataBase.C(TOURN_COLLECTION)

	if db.PronoColl == nil {
		log.Fatal("Fatal error in instantiating the prono collection")
	}

	if db.TournColl == nil {
		log.Fatal("Fatal error in instantiating the tourn collection")
	}
}

//Disabled
/*
func (db *DataBridge) FindPronoById(id string) (Prono, error) {
	var prono Prono
	if !bson.IsObjectIdHex(id) {
		return prono, errors.New("invalid id")
	}
	err := db.PronoCollFindId(bson.ObjectIdHex(id)).One(&prono)

	return prono, err
}
*/
