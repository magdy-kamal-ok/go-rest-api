package main

import (
	"fmt"
	"github.com/magdy-kamal-ok/go-rest-api/internal/comment"
	"github.com/magdy-kamal-ok/go-rest-api/internal/database"
	trasnportHttp "github.com/magdy-kamal-ok/go-rest-api/internal/trasnsport/http"
	"net/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Start running App")
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
		fmt.Println("Error happens")
		return err
	}
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error running Go app")
	}
}
