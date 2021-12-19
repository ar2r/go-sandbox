package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) NumberHandler(w http.ResponseWriter, r *http.Request) {
	app.log.Info("Входящий запрос NumberHandler")

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	num := vars["number"]
	fmt.Fprintf(w, "Number: %v", num)
}

func (app *application) JsonHandler(w http.ResponseWriter, r *http.Request) {
	app.log.Info("Входящий запрос JsonHandler")

	vars := mux.Vars(r)
	num := vars["number"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"value": %v}`, num)
}
