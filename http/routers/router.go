package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Uri      string
	Method   string
	function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Register(r *mux.Router) {
	rout := listRouters

	for _, route := range rout {
		r.HandleFunc(route.Uri, route.function).Methods(route.Method)
	}
}
