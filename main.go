package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SumirVats2003/fem-go/internal/app"
)

func main() {
	app, err := app.InitApp()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Running app")
  http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal()
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Status is available")
}
