package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"youtube/internal/model"
	"youtube/internal/security"
)

type App struct {
	config  *Config
	service *Service
}

func NewApp(config *Config, service *Service) *App {
	return &App{config: config, service: service}
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World!")
	})

	mux.HandleFunc(
		"/api",
		security.ApiKeyMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				u := r.Context().Value("user")
				fmt.Fprintf(w, "Hello world: %s", u.(*model.User).Username)
			},
			a.service.Database,
		),
	)

	log.Printf("Server listening http://localhost:%d\n", a.config.Port)
	return http.ListenAndServe(":"+strconv.Itoa(a.config.Port), mux)
}
