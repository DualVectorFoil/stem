package handler

import (
	"github.com/DualVectorFoil/stem/app/ctrl"
	"github.com/DualVectorFoil/stem/controller"
	"sync"
)

type handler struct {
	UserCtrl *controller.UserCtrl
}

var h *handler
var handlerOnce sync.Once

func NewHandlerInstance() *handler {
	handlerOnce.Do(func() {
		h = &handler{
			UserCtrl: ctrl.UserCtrl,
		}
	})
	return h
}
