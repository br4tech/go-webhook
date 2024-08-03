package adapter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/br4tech/go-webhook/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter[T any] struct {
	client *mongo.Client
	cfg    *config.Config
}

func NewMongoAdapter[T any](
	cfg *config.Config,
) *MongoAdapter[T] {
	queryString := fmt.Sprintf("mongodb://%s:%s@%s:%d/",
		cfg.Db.Mongo.Username, cfg.Db.Mongo.Password, cfg.Db.Mongo.Host, cfg.Db.Mongo.Port)

	clientOptions := options.Client().ApplyURI(queryString)

	// Conectando ao banco de dados
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Ping o banco de dados para verificar a conexão
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return &MongoAdapter[T]{
		cfg:    cfg,
		client: client,
	}
}

func (adapter *MongoAdapter[T]) ClientMongo() *mongo.Client {
	return adapter.client
}

func (adapter *MongoAdapter[T]) FindBy(filter interface{}) ([]T, error) {
	if adapter.client == nil {
		return nil, errors.New("Client is not initialized")
	}

	collection := adapter.client.Database(adapter.cfg.Db.Mongo.DatabaseName).Collection(adapter.cfg.Db.Mongo.CollectionName)

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

	collection := adapter.client.Database(adapter.cfg.Db.Mongo.DatabaseName).Collection(adapter.cfg.Db.Mongo.CollectionName)

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

	collection := adapter.client.Database(adapter.cfg.Db.Mongo.DatabaseName).Collection(adapter.cfg.Db.Mongo.CollectionName)

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
