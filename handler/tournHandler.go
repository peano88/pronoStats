package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peano88/pronoStats/dataLayer"
)

func (hb *HandlerBridge) AddTournament(w http.ResponseWriter, r *http.Request) {
	var t dataLayer.Tournament
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
		return
	}

	newId, err := hb.db.AddTournament(t)
	if err != nil {
		hb.rnd.JSON(w, http.StatusInternalServerError, err.Error)
		return
	}

	hb.rnd.JSON(w, http.StatusOK, newId)

}

func (hb *HandlerBridge) GetTournament(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]

	if !ok {
		hb.rnd.JSON(w, http.StatusBadRequest, "No Id provided")
		return
	}

	t, err := hb.db.FindTournamentById(id)

	if err != nil {
		hb.rnd.JSON(w, http.StatusBadRequest, err.Error)
		return
	}

	hb.rnd.JSON(w, http.StatusOK, t)
}
