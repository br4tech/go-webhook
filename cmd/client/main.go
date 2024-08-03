package main

import (
	"github.com/br4tech/go-webhook/cmd/client/server"
	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/adapter"
	"github.com/br4tech/go-webhook/internal/core/usecase"
	"github.com/br4tech/go-webhook/internal/handler"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
	"github.com/br4tech/go-webhook/internal/repositories"
)

func main() {
	// 1. Carregar a configuração
	cfg := config.GetConfig()

	// 2. Criar o adaptador MongoDB
	productAdapter := adapter.NewPostgresAdapter[model.Product](&cfg)
	orderAdapter := adapter.NewPostgresAdapter[model.Order](&cfg)
	orderItemAdapter := adapter.NewPostgresAdapter[model.OrderItem](&cfg)

	// 3. Criar o repositório
	productRepository := repositories.NewProductRepository(productAdapter)
	orderRepository := repositories.NewOrderRepository(orderAdapter)
	orderItemRepository := repositories.NewOrderItemRepository(orderItemAdapter)

	// 4. Criar o caso de uso
	productUseCase := usecase.NewProductUseCase(productRepository)
	orderUsecase := usecase.NewOrderUseCase(orderRepository, orderItemRepository)

	// 5. Criar o handler
	productHandler := handler.NewProductHandler(productUseCase)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	// 6. Criar o servidor Echo
	echoServer := server.NewEchoServer(&cfg, productAdapter.Db, productHandler, orderHandler)
	// 7. Iniciar o servidor
	echoServer.Start()
}
