package port

import "go.mongodb.org/mongo-driver/mongo"

type IMongoDatabase[T any] interface {
	ClientMongo() *mongo.Client
	FindBy(filter interface{}) ([]T, error)
	FindAll() ([]T, error)
	Create(entity T) error
	Disconnect() error
}
