package main

import (
	"encoding/json"
	"fmt"
	"github.com/ar2r/go-sandbox/cmd/sandbox/dto"
	"github.com/gorilla/mux"
	"math/rand"
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

func (app *application) CreateHandler(w http.ResponseWriter, r *http.Request) {
	app.log.Info("Входящий запрос CreateHandler")

	var req dto.CreatePasteBinRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := dto.PasteBinDto{Id: rand.Uint64(), Title: req.Title, Content: req.Content}
	json.NewEncoder(w).Encode(res)
}

func (app *application) GetHandler(w http.ResponseWriter, r *http.Request) {
	app.log.Info("Входящий запрос GetHandler")

	var req dto.GetPasteBinRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := dto.PasteBinDto{Id: req.Id, Title: "Привет", Content: "Пластимассовый мир"}
	json.NewEncoder(w).Encode(res)
}
