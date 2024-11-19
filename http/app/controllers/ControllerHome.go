package controllers

import (
	"net/http"
	Services "servidorhttp/app/services"
)

type ControllerHome struct {
	Controller
}

func (h ControllerHome) Welcome(response http.ResponseWriter, request *http.Request) {
	sh := Services.ServiceHome{}
	welcome := sh.Welcome()
	h.responseJson(response, welcome)
}
