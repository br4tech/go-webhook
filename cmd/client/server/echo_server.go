package server

import (
	"fmt"

	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app          *echo.Echo
	db           *gorm.DB
	cfg          *config.Config
	orderHandler port.IOrderHandler
}

func NewEchoServer(
	cfg *config.Config,
	db *gorm.DB,
	orderHandler port.IOrderHandler) IServer {

	return &echoServer{
		app:          echo.New(),
		cfg:          cfg,
		orderHandler: orderHandler,
	}
}

// Start inicia o servidor Echo.
func (s *echoServer) Start() {

	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	// s.app.POST("/subscriber", s.webhookHandler)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
