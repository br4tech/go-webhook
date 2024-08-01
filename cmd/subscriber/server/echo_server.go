package server

import (
	"fmt"

	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type echoServer struct {
	app              *echo.Echo
	db               *mongo.Client
	cfg              *config.Config
	subscribeHandler port.ISubscribeHandler
}

func NewEchoServer(
	cfg *config.Config,
	mongo *mongo.Client,
	subscribeHandler port.ISubscribeHandler) IServer {

	return &echoServer{
		app:              echo.New(),
		cfg:              cfg,
		subscribeHandler: subscribeHandler,
	}
}

// Start inicia o servidor Echo.
func (s *echoServer) Start() {

	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/subscriber", s.subscribeHandler.Create)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
