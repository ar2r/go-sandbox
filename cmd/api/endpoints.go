package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет мир!")
	})
	router.HandleFunc("/number/{number:[0-9]+}", NumberHandler)
	router.HandleFunc("/json/{number:[0-9]+}", JsonHandler)

	fmt.Println("Server running")
	srv := &http.Server{
		Handler: router,
		Addr:    ":80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

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
