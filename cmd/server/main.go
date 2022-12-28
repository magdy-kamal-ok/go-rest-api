package main

import (
	"fmt"
	"net/http"

	trasnportHttp "github.com/magdy-kamal-ok/go-rest-api/internal/trasnsport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Start running App")
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
