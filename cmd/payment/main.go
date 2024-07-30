package main

import (
	"github.com/br4tech/go-webhook/cmd/payment/server"
	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/adapter"
	"github.com/br4tech/go-webhook/internal/core/usecase"
	"github.com/br4tech/go-webhook/internal/handler"
	"github.com/br4tech/go-webhook/internal/model"
	"github.com/br4tech/go-webhook/internal/repositories"
)

func main() {
	// 1. Carregar a configuração
	cfg := config.GetConfig()

	// 2. Criar o adaptador MongoDB
	mongoAdapter := adapter.NewMongoAdapter[model.Payment](&cfg)

	// 3. Criar o repositório
	paymentRepository := repositories.NewPaymentRepository(mongoAdapter)

	// 4. Criar o caso de uso
	paymentUseCase := usecase.NewPaymentUsecase(paymentRepository)

	// 5. Criar o handler
	paymentHandler := handler.NewPaymnentHandler(paymentUseCase)

	// 6. Criar o servidor Echo
	echoServer := server.NewEchoServer(&cfg, mongoAdapter.ClientMongo(), paymentHandler)

	// 7. Iniciar o servidor
	echoServer.Start()
}
