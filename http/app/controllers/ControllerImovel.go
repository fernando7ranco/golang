package controllers

import (
	"net/http"
	"servidorhttp/app/services"
	Services "servidorhttp/app/services"
	"strconv"

	"github.com/gorilla/mux"
)

type ControllerImovel struct {
	Controller
	service Services.ServiceImovel
}

func NewControllerImovel() *ControllerImovel {
	service := services.NewServiceImovel()
	return &ControllerImovel{service: service}
}

func (r *ControllerImovel) All(response http.ResponseWriter, request *http.Request) {
	imovels := r.service.All()
	r.responseJson(response, imovels)
}

func (r *ControllerImovel) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idImovel, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(response, "invalid id imovel", http.StatusBadRequest)
		return
	}
	imovel := r.service.Get(uint(idImovel))
	r.responseJson(response, imovel)
}
