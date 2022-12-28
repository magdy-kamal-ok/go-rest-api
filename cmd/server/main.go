package main

import (
	"fmt"
	"net/http"

	db "github.com/magdy-kamal-ok/go-rest-api/internal/databse"
	trasnportHttp "github.com/magdy-kamal-ok/go-rest-api/internal/trasnsport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Start running App")
	var err error
	_, err := db.NewDatabase()
	if err != nil {
		return err
	}
	handler := trasnportHttp.NewHandler()
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
