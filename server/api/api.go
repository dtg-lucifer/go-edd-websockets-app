package api

import (
	"fmt"

	"github.com/dtg-lucifer/go-eda-websocket-app/handlers"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Engine *gin.Engine

	// TODO: storage
}

func NewApi() *Api {
	return &Api{
		Engine: gin.New(),
		// TODO: attach storage instance
	}
}

func (a *Api) Init(version string) error {
	// -----------------------
	// Initialise the API
	a.Engine.Use(gin.Recovery())
	a.Engine.Use(gin.Logger())

	a.SetupRoutes(version)

	// -----------------------
	// Initialise the Storage
	// TODO: make migrations

	return nil
}

func (a *Api) SetupRoutes(version string) {
	// Health Route
	a.Engine.GET(fmt.Sprintf("%s/health", version), handlers.HealthCheckRoute)
}

func (a *Api) Start(addr string) error {
	return a.Engine.Run(addr)
}
