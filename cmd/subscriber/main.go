package main

import (
	"github.com/br4tech/go-webhook/cmd/subscriber/server"
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
	mongoAdapter := adapter.NewMongoAdapter[model.Subscribe](&cfg)

	// 3. Criar o repositório
	subscribeRepository := repositories.NewSubscribeRespository(mongoAdapter)

	// 4. Criar o caso de uso
	subscribeUseCase := usecase.NewSubscribeUseCase(subscribeRepository)

	// 5. Criar o handler
	subscribeHandler := handler.NewSubscribeHandler(subscribeUseCase)

	// 6. Criar o servidor Echo
	echoServer := server.NewEchoServer(&cfg, mongoAdapter.ClientMongo(), subscribeHandler)

	// 7. Iniciar o servidor
	echoServer.Start()
}
