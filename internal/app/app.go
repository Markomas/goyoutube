package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Config *Config
}

func NewApp(config *Config) *App {
	return &App{Config: config}
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World!")
	})

	log.Printf("Server listening http://localhost:%d\n", a.Config.Port)
	return http.ListenAndServe(":"+strconv.Itoa(a.Config.Port), mux)
}
