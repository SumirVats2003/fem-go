package routes

import (
	"github.com/SumirVats2003/fem-go/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	// health check
	r.Get("/health", app.HealthCheck)

	// workout routes
	r.Get("/workout/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workout/create", app.WorkoutHandler.HandleCreateWorkout)
	return r
}
