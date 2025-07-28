package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SumirVats2003/fem-go/internal/api"
	"github.com/SumirVats2003/fem-go/internal/store"
	"github.com/SumirVats2003/fem-go/migrations"
)

type App struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func InitApp() (*App, error) {
	postgres, err := store.ConnectDB()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(postgres, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	app := &App{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             postgres,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
