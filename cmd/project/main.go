package main

import (
	"../../adapter/rest"
	"../../app"
	"../../models/dbs"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AutomaticEnv()

	bindAddress := viper.GetString("bind_address")
	if bindAddress == "" {
		bindAddress = "localhost:7481"
	}

	projectService := rest.New(&rest.Config{
		Version:     viper.GetString("version"),
		BindAddress: bindAddress,
		AppConfig:   &app.Config{
			DBInfo: &dbs.DBConfig{
				PostgresHost:     viper.GetString("postgres_host"),
				PostgresPort:     viper.GetInt("postgres_port"),
				PostgresUser:     viper.GetString("postgres_user"),
				PostgresDBName:   viper.GetString("postgres_db_name"),
				PostgresPassword: viper.GetString("postgres_password"),
			},
		},
	})

	log.Fatal(projectService.Run())
}


