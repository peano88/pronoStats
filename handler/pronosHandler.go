package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peano88/pronoStats/dataLayer"
)

type PronosResponse struct {
	Pronos     dataLayer.TournamentPronos
	Tournament dataLayer.Tournament
}

func (hb *HandlerBridge) AddTournamentPronos(w http.ResponseWriter, r *http.Request) {
	var tPronos dataLayer.TournamentPronos
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&tPronos); err != nil {
		hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
		return
	}

	newId, err := hb.db.AddTournamentPronos(tPronos)
	if err != nil {
		hb.rnd.JSON(w, http.StatusInternalServerError, err.Error)
		return
	}

	hb.rnd.JSON(w, http.StatusOK, newId)

}

func (hb *HandlerBridge) GetTournamentPronos(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]

	if !ok {
		hb.rnd.JSON(w, http.StatusBadRequest, "No Id provided")
		return
	}

	var pr PronosResponse
	var err error
	pr.Pronos, err = hb.db.FindTournamentPronosById(id)

	if err != nil {
		hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
		return
	}
	if pr.Pronos.TournamentId != "" {
		pr.Tournament, err = hb.db.FindTournamentById(pr.Pronos.TournamentId)
		if err != nil {
			hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
			return
		}
	}

	hb.rnd.JSON(w, http.StatusOK, pr)
}

func (hb *HandlerBridge) GetTournamentPronosByUser(w http.ResponseWriter, r *http.Request) {
	user, ok := mux.Vars(r)["user"]

	if !ok {
		hb.rnd.JSON(w, http.StatusBadRequest, "No user provided")
		return
	}

	tPronos, err := hb.db.FindTournamentPronosByUser(user)

	if err != nil {
		hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
		return
	}

	tP := make([]PronosResponse, len(tPronos))
	for _, p := range tPronos {
		var tourn dataLayer.Tournament
		if p.TournamentId != "" {
			tourn, err = hb.db.FindTournamentById(p.TournamentId)
			if err != nil {
				hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
				return
			}
		}
		aTournPronos := PronosResponse{
			Pronos:     p,
			Tournament: tourn,
		}

		tP = append(tP, aTournPronos)
	}

	hb.rnd.JSON(w, http.StatusOK, tP)
}

func (hb *HandlerBridge) AddProno(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id_tp"]
	if !ok {
		hb.rnd.JSON(w, http.StatusBadRequest, "No Id provided")
		return
	}

	var prono dataLayer.Prono
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&prono); err != nil {
		hb.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := hb.db.AddProno(id, prono)
	if err != nil {
		hb.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	hb.rnd.JSON(w, http.StatusOK, "")

}

//Disabled
/*
func (hb *HandlerBridge) GetProno(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]

	if !ok {
		hb.rnd.JSON(w, http.StatusBadRequest, "No Id provided")
		return
	}

	prono, err := hb.db.FindPronoById(id)

	if err != nil {
		hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
		return
	}

	hb.rnd.JSON(w, http.StatusOK, prono)

}
*/
