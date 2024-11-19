package controllers

import (
	"net/http"
	interfaces "servidorhttp/app/interfaces"
)

type Controller struct{}

func (c Controller) responseJson(response http.ResponseWriter, obj interfaces.IStructToJson) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(obj.ToJson())
}
