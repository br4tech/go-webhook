package port

type IMongoDatabase[T any] interface {
	FindBy(filter interface{}) ([]T, error)
	FindAll() ([]T, error)
	Create(entity T) error
	Disconnect() error
}
