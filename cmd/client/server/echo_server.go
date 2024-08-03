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
	app            *echo.Echo
	cfg            *config.Config
	orderHandler   port.IOrderHandler
	productHandler port.IProductHandler
}

func NewEchoServer(
	cfg *config.Config,
	db *gorm.DB,
	orderHandler port.IOrderHandler,
	productHandler port.IProductHandler,
) IServer {
	return &echoServer{
		app:            echo.New(),
		cfg:            cfg,
		orderHandler:   orderHandler,
		productHandler: productHandler,
	}
}

// Start inicia o servidor Echo.
func (s *echoServer) Start() {

	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/product", s.productHandler.Create)
	s.app.POST("/order", s.orderHandler.Create)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.ClientPort)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
