package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guiyones/fcutils/pkg/rabbitmq"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	err := os.Chdir("..") // volta um nível
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	queue, err := rabbitmq.QueueDeclare(ch)
	if err != nil {
		panic(err)
	}

	go rabbitmq.Consume(ch, msgs, queue)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // Para não voltar a mensagem para fila
	}
}
