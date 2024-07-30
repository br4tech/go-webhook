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
	app            *echo.Echo
	db             *mongo.Client
	cfg            *config.Config
	paymentHandler port.IPaymentHandler
}

func NewEchoServer(
	cfg *config.Config,
	mongo *mongo.Client,
	paymentHandler port.IPaymentHandler) IServer {

	return &echoServer{
		app:            echo.New(),
		cfg:            cfg,
		paymentHandler: paymentHandler,
	}
}

// Start inicia o servidor Echo.
func (s *echoServer) Start() {

	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/payment", s.paymentHandler.Create)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
