package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет мир!")
	})
	r.HandleFunc("/double/{number:[0-9]+}", DoubleHandler)

	fmt.Println("Server running")
	http.ListenAndServe(":80", r)
}

func DoubleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	num := vars["number"]
	fmt.Fprintf(w, "Number wow: %v\n", num)
}
