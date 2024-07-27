package port

type IMongoDatabase[T any] interface {
	Connect(uri string) error
	FindBy(filter interface{}) ([]T, error)
	FindAll() ([]T, error)
	Create(entity *T) error
	Disconnect() error
}
