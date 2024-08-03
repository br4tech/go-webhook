package main

import (
	"github.com/br4tech/go-webhook/cmd/client/server"
	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/adapter"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/core/usecase"
	"github.com/br4tech/go-webhook/internal/handler"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
	"github.com/br4tech/go-webhook/internal/repositories"
)

func main() {
	// 1. Carregar a configuração
	cfg := config.GetConfig()

	// 2. Criar o adaptador MongoDB
	postgresAdapter := adapter.NewPostgresAdapter[port.IModel](&cfg)

	// 3. Criar o repositório
	productRepository := repositories.NewProductRepository(postgresAdapter.(port.IPostgreDatabase[model.Product]))
	orderRepository := repositories.NewOrderRepository(postgresAdapter.(port.IPostgreDatabase[model.Order]))
	orderItemRepository := repositories.NewOrderItemRepository(postgresAdapter.(port.IPostgreDatabase[model.OrderItem]))

	// 4. Criar o caso de uso
	productUseCase := usecase.NewProductUseCase(productRepository)
	orderUsecase := usecase.NewOrderUseCase(orderRepository, orderItemRepository)

	// 5. Criar o handler
	productHandler := handler.NewProductHandler(productUseCase)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	// 6. Criar o servidor Echo
	echoServer := server.NewEchoServer(&cfg, postgresAdapter.GetDb(), productHandler, orderHandler)
	// 7. Iniciar o servidor
	echoServer.Start()
}
