package port

import "go.mongodb.org/mongo-driver/mongo"

type IMongoDatabase[T any] interface {
	ClientMongo() *mongo.Client
	FindBy(filter interface{}) ([]T, error)
	FindAll() ([]T, error)
	Create(entity T) error
	Disconnect() error
}

type IPostgreDatabase[T any] interface {
	FindAll() ([]T, error)
	Create(entity T) (int, error)
}

type IModel interface {
	GetId() int
}
