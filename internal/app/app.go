package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	Logger *log.Logger
}

func InitApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &App{
		Logger: logger,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
