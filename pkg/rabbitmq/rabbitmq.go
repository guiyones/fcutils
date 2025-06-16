package rabbitmq

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Cria uma conexão e um canal para trabalharmos dentro da conexão
func OpenChannel() (*amqp.Channel, error) {
	// Cria a conexão
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// Esse é um canal do RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func QueueDeclare(ch *amqp.Channel) (amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		os.Getenv("RABBITMQ_QUEUE_NAME"),
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return amqp.Queue{}, err
	}

	return queue, nil
}

// Consumir mensagem do RabbitMQ que esta em uma fila
// O parametro out pega o resultado da mensagem que foi recebida pelo consumo
func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue amqp.Queue) error {
	msgs, err := ch.Consume(
		queue.Name,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exName string, contentType string, bindKey string) error {
	err := ch.Publish(
		exName,
		bindKey,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}
	return nil
}
