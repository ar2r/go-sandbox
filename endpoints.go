package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NumberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	num := vars["number"]
	fmt.Fprintf(w, "Number: %v", num)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["number"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"value": %v}`, num)
}
