package server

import (
	"fmt"

	"github.com/br4tech/go-webhook/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type EchoServer struct {
	app *echo.Echo
	db  *mongo.Client
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *mongo.Client) IServer {
	return &EchoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *EchoServer) Start() {
	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
