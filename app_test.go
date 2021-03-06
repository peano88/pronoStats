package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/peano88/pronoStats/dataLayer"
	"github.com/stretchr/testify/assert"
)

type requesTest struct {
	endpoint   string
	method     string
	prono      dataLayer.Prono
	tournament dataLayer.TournamentPronos
	expStatus  int
	label      string
}

var baseProno dataLayer.Prono = dataLayer.Prono{
	HomeTeam:       "TeamA",
	AwayTeam:       "TeamB",
	HomeScore:      2,
	AwayScore:      1,
	PronoAwayScore: 2,
	PronoHomeScore: 2,
}

var baseTournament = dataLayer.TournamentPronos{
	Tournament: "Tourn 1",
	Sport:      "Sport Z",
	User:       "UserTest",
	Pronos:     []dataLayer.Prono{baseProno},
}

const (
	BASE_ENDPOINT_TOURNAMENT = "/tournaments"
)

func Test_Endpoints(t *testing.T) {
	var requests []requesTest

	requests = append(requests, requesTest{
		endpoint:   BASE_ENDPOINT_TOURNAMENT,
		method:     "POST",
		tournament: baseTournament,
		expStatus:  http.StatusOK,
		label:      "basic creation without id",
	})

	tournamentWithId := baseTournament
	tournamentWithId.ID = bson.NewObjectId()
	requests = append(requests, requesTest{
		endpoint:   BASE_ENDPOINT_TOURNAMENT,
		method:     "POST",
		tournament: tournamentWithId,
		expStatus:  http.StatusOK,
		label:      "basic creation with id",
	})

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT_TOURNAMENT + "/" + tournamentWithId.ID.Hex(),
		method:    "GET",
		expStatus: http.StatusOK,
		label:     "basic get",
	})

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT_TOURNAMENT + "/010102345687123456789012",
		method:    "GET",
		expStatus: http.StatusBadRequest,
		label:     "get with non-existing id",
	})

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT_TOURNAMENT + "/" + tournamentWithId.ID.Hex() + "/prono",
		method:    "POST",
		prono:     baseProno,
		expStatus: http.StatusOK,
		label:     "Create prono",
	})

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT_TOURNAMENT + "?user=" + baseTournament.User,
		method:    "GET",
		expStatus: http.StatusOK,
		label:     "get tprono by user",
	})

	for i, rt := range requests {
		buf := new(bytes.Buffer)
		if rt.tournament.Tournament != "" {
			if err := json.NewEncoder(buf).Encode(rt.tournament); err != nil {
				t.Fatalf("Test %d error : %s", i, err.Error())
			}
		}
		if rt.prono.HomeTeam != "" {
			if err := json.NewEncoder(buf).Encode(rt.prono); err != nil {
				t.Fatalf("Test %d error : %s", i, err.Error())
			}
		}
		req, _ := http.NewRequest(rt.method, rt.endpoint, buf)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, rt.expStatus, rr.Code, rt.label)

		/*
			if rt.expProno.HomeTeam != "" {
				var received dataLayer.Prono
				json.NewDecoder(rr.Body).Decode(&received)
				if assert.NotEmpty(t, received) {
					assert.Equal(t, rt.expProno, received)
				}
			}
		*/

	}
}
