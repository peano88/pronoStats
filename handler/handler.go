package handler

import (
	"net/http"

	"github.com/peano88/pronoStats/dataLayer"
)

type HandlerBridge struct {
	Db dataLayer.DataBridge
}

func (hb *HandlerBridge) AddProno(w http.ResponseWriter, r *http.Request) {
}

func (hb *HandlerBridge) GetProno(w http.ResponseWriter, r *http.Request) {
}
