package users

import (
	"../internal/constant"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Config struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

type UserDatabase struct {
	table *gorm.DB
}

func Open(config *Config) (*UserDatabase, error) {
	table, err := gorm.Open("postgres", fmt.Sprintf(constant.PostgresConnStr, config.Host, config.Port, config.User, config.DBName, config.Password))
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf(constant.ConnectingError + constant.PostgresConnStr, config.Host, config.Port, config.User, config.DBName, config.Password))
	}

	return &UserDatabase{
		table: table,
	}, nil
}

func (schema *UserDatabase) Initialize() {
	schema.table.AutoMigrate(&User{})
}

func (schema *UserDatabase) Close() error {
	return schema.table.Close()
}