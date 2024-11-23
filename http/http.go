package main

import (
	"log"
	"net/http"
	"servidorhttp/routers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.Register(r)
	log.Fatal(http.ListenAndServe(":666", r))
}
