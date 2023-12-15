package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"youtube/internal/app"
	"youtube/internal/config"
	"youtube/internal/db"
)

func main() {
	//TODO: config options
	migrate := true
	populate := true

	//load config file
	conf, err := config.LoadConfig("config.yaml")
	fmt.Printf("%+v\n", conf)

	//create database connection
	gormDb, err := gorm.Open(sqlite.Open(conf.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database := db.NewDatabase(gormDb)
	if migrate {
		err = database.Migrate()
	}

	if populate {
		err = database.Populate()
	}

	//app build
	service := app.Service{
		Database: database,
	}

	server := app.NewApp(&conf.App, &service)
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
