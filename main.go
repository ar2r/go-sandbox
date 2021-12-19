package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// Структура для хранения зависимостей
type application struct {
	logger *log.Logger
}

func main() {
	// Добавляем json логирование
	myLogger := log.New()
	myLogger.SetFormatter(&log.JSONFormatter{})
	myLogger.SetOutput(os.Stdout)
	myLogger.SetLevel(log.DebugLevel)

	// Конфигурация приложения
	app := &application{
		logger: myLogger,
	}

	// Пример логирования
	app.logger.Error("Error Message")
	app.logger.Warning("Warning Message")
	app.logger.Info("Info message")
	app.logger.Debug("Debug Message")
	app.logger.Trace("Trace message")

	app.logger.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет мир!")
	})
	router.HandleFunc("/number/{number:[0-9]+}", app.NumberHandler)
	router.HandleFunc("/json/{number:[0-9]+}", app.JsonHandler)
	app.logger.Info("Routes registered")

	srv := &http.Server{
		Handler: router,
		Addr:    ":80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	app.logger.Info("Server running. Waiting for income requests.")
	app.logger.Fatal(srv.ListenAndServe())
}
