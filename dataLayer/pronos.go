package dataLayer

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type Prono struct {
	HomeTeam        string `bson:"home_team" json:"home_team"`
	AwayTeam        string `bson:"away_team" json:"away_team"`
	HomeScore       int    `bson:"home_score" json:"home_score"`
	AwayScore       int    `bson:"away_score" json:"away_score"`
	PronoHomeScore  int    `bson:"prono_home_score" json:"prono_home_score"`
	PronoAwayScore  int    `bson:"prono_away_score" json:"prono_away_score"`
	PronoDifference int    `bson:"prono_diff" json:"prono_diff"`
}

type TournamentPronos struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	User       string        `bson:"user" json:"user"`
	Sport      string        `bson:"sport" json:"sport"`
	Tournament string        `bson:"tournament" json:"tournament"`
	PronoDiff  bool          `bson:"prono_diff" json:"prono_diff"`
	Pronos     []Prono       `bson:"pronos" json:"pronos"`
}

func (db *DataBridge) AddTournamentPronos(tp TournamentPronos) (string, error) {
	if tp.ID == "" {
		tp.ID = bson.NewObjectId()
	}

	if err := db.PronoColl.Insert(&tp); err != nil {
		return "", err
	}

	return tp.ID.Hex(), nil
}

func (db *DataBridge) FindTournamentPronosById(id string) (TournamentPronos, error) {
	var tPronos TournamentPronos

	if !bson.IsObjectIdHex(id) {
		return tPronos, errors.New("invalid id")
	}
	err := db.PronoColl.FindId(bson.ObjectIdHex(id)).One(&tPronos)

	return tPronos, err
}

func (db *DataBridge) FindTournamentPronosByUser(user string) ([]TournamentPronos, error) {
	var tPronos []TournamentPronos

	err := db.PronoColl.Find(bson.M{"user": user}).All(&tPronos)

	return tPronos, err
}

func (db *DataBridge) AddProno(idTourPronos string, pr Prono) error {
	if !bson.IsObjectIdHex(idTourPronos) {
		return errors.New("invalid id")
	}
	return db.PronoColl.UpdateId(bson.ObjectIdHex(idTourPronos), bson.M{"$push": bson.M{"pronos": pr}})

}
