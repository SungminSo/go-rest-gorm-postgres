package app

import (
	"../models/admins"
	"../models/dbs"
	"../models/users"
	"log"
)

type Config struct {
	DBInfo *dbs.DBConfig
}

type ProjectApp struct {
	config *Config

	admins *admins.AdminDatabase
	users  *users.UserDatabase
}

func New(config *Config) *ProjectApp {
	return &ProjectApp{
		config: config,
	}
}

func (app *ProjectApp) Initialize() error {
	var err error

	app.admins, err = admins.Open(&admins.Config{
		Host:     app.config.DBInfo.PostgresHost,
		Port:     app.config.DBInfo.PostgresPort,
		User:     app.config.DBInfo.PostgresUser,
		DBName:   app.config.DBInfo.PostgresDBName,
		Password: app.config.DBInfo.PostgresPassword,
	})
	if err != nil {
		return err
	}
	app.admins.Initialize()

	app.users, err = users.Open(&users.Config{
		Host:     app.config.DBInfo.PostgresHost,
		Port:     app.config.DBInfo.PostgresPort,
		User:     app.config.DBInfo.PostgresUser,
		DBName:   app.config.DBInfo.PostgresDBName,
		Password: app.config.DBInfo.PostgresPassword,
	})
	if err != nil {
		return err
	}
	app.users.Initialize()

	return nil
}

func (app *ProjectApp) Close() {
	err := app.admins.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = app.users.Close()
	if err != nil {
		log.Fatal(err)
	}
}
