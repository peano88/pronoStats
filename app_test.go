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
	endpoint  string
	method    string
	prono     dataLayer.Prono
	expStatus int
	expProno  dataLayer.Prono
}

var baseProno dataLayer.Prono = dataLayer.Prono{
	HomeTeam:       "TeamA",
	AwayTeam:       "TeamB",
	HomeScore:      2,
	AwayScore:      1,
	PronoAwayScore: 2,
	PronoHomeScore: 2,
}

const (
	BASE_ENDPOINT = "/"
)

func Test_Endpoints(t *testing.T) {
	var requests []requesTest

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT,
		method:    "POST",
		prono:     baseProno,
		expStatus: http.StatusOK,
	})

	pronoWithId := baseProno
	pronoWithId.ID = bson.NewObjectId()
	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT,
		method:    "POST",
		prono:     pronoWithId,
		expStatus: http.StatusOK,
	})

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT + pronoWithId.ID.Hex(),
		method:    "GET",
		expStatus: http.StatusOK,
		expProno:  pronoWithId,
	})

	requests = append(requests, requesTest{
		endpoint:  BASE_ENDPOINT + "not existing",
		method:    "GET",
		expStatus: http.StatusBadRequest,
	})

	for i, rt := range requests {
		buf := new(bytes.Buffer)
		if rt.prono.HomeTeam != "" {
			if err := json.NewEncoder(buf).Encode(rt.prono); err != nil {
				t.Fatalf("Test %d error : %s", i, err.Error())
			}
		}
		req, _ := http.NewRequest(rt.method, rt.endpoint, buf)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, rt.expStatus, rr.Code)
		if rt.expProno.HomeTeam != "" {
			var received dataLayer.Prono
			json.NewDecoder(rr.Body).Decode(&received)
			if assert.NotEmpty(t, received) {
				assert.Equal(t, rt.expProno, received)
			}
		}

	}
}
