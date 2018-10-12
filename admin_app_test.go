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

type reqTourn struct {
	endpoint   string
	method     string
	tournament dataLayer.Tournament
	expStatus  int
}

const (
	BASE_ENDPOINT_TOURNAMENT_ADMIN = "/admin/tournaments"
)

func Test_Tournaments(t *testing.T) {
	var requests []reqTourn

	id := bson.NewObjectId()

	newTournament := dataLayer.Tournament{
		ID:    id,
		Sport: "Footy",
		Name:  "Champion's League",
		Teams: []dataLayer.Team{
			dataLayer.Team{
				Name: "Crying Babies",
				Categories: []string{
					"Simulators",
					"Somewhere",
				},
			},
			dataLayer.Team{
				Name: "Pussies",
				Categories: []string{
					"Simulators++",
					"Somewhere else",
				},
			},
		},
		Rounds: []dataLayer.Round{
			dataLayer.Round{
				Number: 1,
				Matches: []dataLayer.Match{
					dataLayer.Match{
						HomeTeam:  "Crying Babies",
						AwayTeam:  "Pussies",
						HomeScore: 2,
						AwayScore: 2,
					},
				},
			},
		},
	}

	requests = append(requests, reqTourn{
		endpoint:   BASE_ENDPOINT_TOURNAMENT_ADMIN,
		method:     "POST",
		tournament: newTournament,
		expStatus:  http.StatusOK,
	})
	requests = append(requests, reqTourn{
		endpoint:   BASE_ENDPOINT_TOURNAMENT_ADMIN + "/" + id.Hex(),
		method:     "GET",
		tournament: dataLayer.Tournament{},
		expStatus:  http.StatusOK,
	})

	for i, rt := range requests {
		buf := new(bytes.Buffer)
		if rt.tournament.Name != "" {
			if err := json.NewEncoder(buf).Encode(rt.tournament); err != nil {
				t.Fatalf("Test %d error : %s", i, err.Error())
			}
		}
		req, _ := http.NewRequest(rt.method, rt.endpoint, buf)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, rt.expStatus, rr.Code)

	}
}
