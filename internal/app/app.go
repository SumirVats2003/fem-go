package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SumirVats2003/fem-go/internal/api"
)

type App struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func InitApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	app := &App{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
