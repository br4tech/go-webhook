package adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/br4tech/go-webhook/config"
	"github.com/br4tech/go-webhook/internal/core/port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter[T any] struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}

func NewMongoAdapter[T any](
	cfg *config.Config,
	databaseName, dollectionName string,
) port.IMongoDatabase[T] {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil
	}

	return &MongoAdapter[T]{
		Client: client,
	}
}

func (adapter *MongoAdapter[T]) FindBy(filter interface{}) ([]T, error) {
	if adapter.Client == nil {
		return nil, errors.New("Client is not initialized")
	}

	collection := adapter.Client.Database(adapter.DatabaseName).Collection(adapter.CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (adapter *MongoAdapter[T]) FindAll() ([]T, error) {
	if adapter.Client == nil {
		return nil, errors.New("client is not connected")
	}

	collection := adapter.Client.Database(adapter.DatabaseName).Collection(adapter.CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find all documents (no filter)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (adapter *MongoAdapter[T]) Create(entity T) error {
	if adapter.Client == nil {
		return errors.New("client is not connected")
	}

	collection := adapter.Client.Database(adapter.DatabaseName).Collection(adapter.CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *MongoAdapter[T]) Disconnect() error {
	if adapter.Client == nil {
		return errors.New("client is not connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := adapter.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	adapter.Client = nil // Limpa a referência ao cliente
	return nil
}
