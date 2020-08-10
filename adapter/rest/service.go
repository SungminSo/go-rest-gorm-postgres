package rest

import (
	"../../app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type Config struct {
	Version 	string
	BindAddress string
	AppConfig	*app.Config
}

type ProjectService struct {
	config 	*Config
	server 	*gin.Engine
	app 	*app.ProjectApp
}

func New(config *Config) *ProjectService {
	projectService := &ProjectService{}
	projectService.Initialize(config)

	return projectService
}

func (ps *ProjectService) Initialize(config *Config) {
	ps.config = config

	ps.app = app.New(config.AppConfig)

	ps.server = gin.New()

	allowOrigins := strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	allowHeaders := strings.Split(os.Getenv("ALLOW_HEADERS"), ",")
	allowMethods := strings.Split(os.Getenv("ALLOW_METHODS"), ",")

	c := cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     allowMethods,
		AllowHeaders:     allowHeaders,
		AllowCredentials: true,
	}

	ps.server.Use(gin.Logger(), gin.Recovery(), cors.New(c))
	ps.registerRoutes()
}

func (ps *ProjectService) Run() error {
	err := ps.app.Initialize()
	if err != nil {
		return err
	}

	return ps.server.Run(ps.config.BindAddress)
}

func (ps *ProjectService) Close() {
	ps.app.Close()
}
