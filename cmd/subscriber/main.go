package main

import (
	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/adapter"
	"github.com/br4tech/go-webhook/internal/core/usecase"
	"github.com/br4tech/go-webhook/internal/handler"
	"github.com/br4tech/go-webhook/internal/model"
	"github.com/br4tech/go-webhook/internal/repositories"
	"github.com/br4tech/go-webhook/server"
)

func main() {
	// 1. Carregar a configuração
	cfg := config.GetConfig()

	// 2. Criar o adaptador MongoDB
	mongoAdapter := adapter.NewMongoAdapter[model.Subscription](&cfg, cfg.Db.DatabaseName, cfg.Db.CollectionName)

	// 3. Criar o repositório
	subscriptionRepository := repositories.NewSubscriptionRespository(mongoAdapter)

	// 4. Criar o caso de uso
	subscriptionUseCase := usecase.NewSubscriptionUseCase(subscriptionRepository)

	// 5. Criar o handler
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionUseCase)

	// 6. Criar o servidor Echo
	echoServer := server.NewEchoServer(&cfg, mongoAdapter, subscriptionHandler)

	// 7. Iniciar o servidor
	echoServer.Start()
}
