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
	log *log.Logger
}

func main() {
	// Добавляем json логирование
	myLog := log.New()
	myLog.SetFormatter(&log.JSONFormatter{})
	myLog.SetOutput(os.Stdout)
	myLog.SetLevel(log.DebugLevel)

	// Конфигурация приложения
	app := &application{
		log: myLog,
	}

	// Пример логирования
	app.log.Error("Error Message")
	app.log.Warning("Warning Message")
	app.log.Info("Info message")
	app.log.Debug("Debug Message")
	app.log.Trace("Trace message")

	app.log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// Регистрация методов API
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет мир!")
	})
	router.HandleFunc("/number/{number:[0-9]+}", app.NumberHandler)
	router.HandleFunc("/json/{number:[0-9]+}", app.JsonHandler)
	app.log.Info("Routes registered")

	// Запуск сервера
	srv := &http.Server{
		Handler: router,
		Addr:    ":80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	app.log.Info("Server running. Waiting for income requests.")
	app.log.Fatal(srv.ListenAndServe())
}
