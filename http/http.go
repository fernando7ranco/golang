package main

import (
	"log"
	"net/http"
	Router "servidorhttp/routers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	Router.Register(r)

	log.Fatal(http.ListenAndServe(":666", r))
}
