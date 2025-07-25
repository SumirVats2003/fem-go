package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/SumirVats2003/fem-go/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.InitApp()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
  app.Logger.Printf("server running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal()
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
