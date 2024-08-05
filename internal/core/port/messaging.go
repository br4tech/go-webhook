package port

import amqp "github.com/rabbitmq/amqp091-go"

type IMessaging interface {
	Publish(queueName, message string) error
	Consume(queueName string, handler func(amqp.Delivery)) error
	Close() error
}
