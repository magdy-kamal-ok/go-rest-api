package main

import (
	"fmt"
	"github.com/magdy-kamal-ok/go-rest-api/internal/comment"
	"github.com/magdy-kamal-ok/go-rest-api/internal/database"
	trasnportHttp "github.com/magdy-kamal-ok/go-rest-api/internal/trasnsport/http"
	"net/http"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields {
			"AppName": app.Name,
			"AppVersion": app.Version,
		}).Info("Start running App")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Error happens", err)
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db)
	handler := trasnportHttp.NewHandler(commentService)
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Error happens")
		return err
	}
	return nil
}

func main() {
	app :=       App{
		Name: "Commenting Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error running Go app")
		log.Fatal(err)
	}
}
