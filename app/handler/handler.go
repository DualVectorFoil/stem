package handler

import (
	"github.com/DualVectorFoil/stem/app/ctrl"
	"github.com/DualVectorFoil/stem/controller"
	"sync"
)

type handler struct {
	UserCtrl  *controller.UserCtrl
	ImageCtrl *controller.ImageCtrl
}

var h *handler
var handlerOnce sync.Once

func NewHandlerInstance() *handler {
	handlerOnce.Do(func() {
		h = &handler{
			UserCtrl:  ctrl.UserCtrl,
			ImageCtrl: ctrl.ImageCtrl,
		}
	})
	return h
}
