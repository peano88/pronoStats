package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peano88/pronoStats/dataLayer"
	"github.com/thedevsaddam/renderer"
)

type HandlerBridge struct {
	db  dataLayer.DataBridge
	rnd *renderer.Render
}

func (hb *HandlerBridge) Init(d dataLayer.DataBridge) {
	hb.db = d
	hb.rnd = renderer.New()
}

func (hb *HandlerBridge) AddProno(w http.ResponseWriter, r *http.Request) {
	var prono dataLayer.Prono
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&prono); err != nil {
		hb.rnd.JSON(w, http.StatusInternalServerError, err.Error)
		return
	}

	newId, err := hb.db.AddProno(prono)
	if err != nil {
		hb.rnd.JSON(w, http.StatusInternalServerError, err.Error)
		return
	}

	hb.rnd.JSON(w, http.StatusOK, newId)

}

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
