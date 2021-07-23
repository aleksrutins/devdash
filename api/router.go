package api

import (
	"github.com/gorilla/mux"
)

func ApiHandler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/listContainers", ListContainers)
	return r
}
