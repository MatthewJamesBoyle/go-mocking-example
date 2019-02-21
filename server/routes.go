package server

import "github.com/gorilla/mux"

func Routes(h *CalcHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/add/", h.AddNumbers).Methods("POST")
	r.HandleFunc("/minus", h.MinusNumbers).Methods("POST")
	return r
}
