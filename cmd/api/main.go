package main

import (
	"youtube/internal/app"
	"youtube/internal/config"
)

func main() {
	conf, err := config.LoadConfig("config.yaml")

	server := app.NewApp(&conf.App)
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
