package controllers

import (
	"net/http"
	"servidorhttp/app/structs"
)

type Controller struct {
	helperJson structs.HelperJson
}

func (r *Controller) responseJson(response http.ResponseWriter, obj interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(r.helperJson.ParseJson(obj))
}
