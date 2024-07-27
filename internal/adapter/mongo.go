package adapter

import (
	"context"
	"errors"
	"time"

	"github.com/br4tech/go-webhook/internal/core/port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter[T any] struct {
	client         *mongo.Client
	uri            string
	databaseName   string
	collectionName string
}

func NewMongoAdapter[T any](uri, databaseName, collectionName string) port.IMongoDatabase[T] {
	return &MongoAdapter[T]{
		uri:            uri,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (adapter *MongoAdapter[T]) Connect(uri string) error {
	if adapter.client != nil {
		return errors.New("client is already connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(adapter.uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	adapter.client = client
	return nil
}

func (adapter *MongoAdapter[T]) FindBy(filter interface{}) ([]T, error) {
	if adapter.client == nil {
		return nil, errors.New("Client is not initialized")
	}

	collection := adapter.client.Database(adapter.databaseName).Collection(adapter.collectionName)

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
	if adapter.client == nil {
		return nil, errors.New("client is not connected")
	}

	collection := adapter.client.Database(adapter.databaseName).Collection(adapter.collectionName)

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
	if adapter.client == nil {
		return errors.New("client is not connected")
	}

	collection := adapter.client.Database(adapter.databaseName).Collection(adapter.collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *MongoAdapter[T]) Disconnect() error {
	if adapter.client == nil {
		return errors.New("client is not connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := adapter.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	adapter.client = nil // Limpa a referência ao cliente
	return nil
}
