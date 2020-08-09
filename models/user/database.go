package user

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/SungminSo/go-rest-gorm-postgres/internal/constant"
	"github.com/SungminSo/go-rest-gorm-postgres/models"
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
	schema.table.AutoMigrate(&models.User{})
}

func (schema *UserDatabase) Close() error {
	return schema.table.Close()
}