package app

import (
	"log"
	"os"
)

type App struct {
  Logger *log.Logger
}

func InitApp() (*App, error) {
  logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

  app := &App {
    Logger: logger,
  }

  return app, nil
}
