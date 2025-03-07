// messaging_service.go
package infrastructure

import (
	"log"

	"github.com/streadway/amqp"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
)

type MessagingService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewMessagingService() (domain.MessagingService, error) {
	conn, err := amqp.Dial("amqp://reyhades:reyhades@44.223.218.9:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return nil, err
	}

	// Declarar el intercambio (opcional, pero recomendado para asegurarse de que exista)
	err = ch.ExchangeDeclare(
		"orders_created", // nombre del intercambio
		"direct",         // tipo de intercambio (direct, topic, fanout, etc.)
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
		return nil, err
	}

	return &MessagingService{conn: conn, ch: ch}, nil
}

func (ms *MessagingService) PublishOrderCreated(order []byte) error {
	// Publicar el mensaje en el intercambio personalizado
	err := ms.ch.Publish(
		"orders_created", // nombre del intercambio
		"order_created",  // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        order,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
		return err
	}
	return nil
}

func (ms *MessagingService) Close() {
	ms.ch.Close()
	ms.conn.Close()
}
