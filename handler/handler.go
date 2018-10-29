package handler

import (
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
