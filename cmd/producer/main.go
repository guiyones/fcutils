package main

import (
	"log"
	"os"

	"github.com/guiyones/fcutils/pkg/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {

	err := os.Chdir("..") // volta um nível
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	exName := os.Getenv("RABBITMQ_EXCHANGE_NAME")
	key := os.Getenv("RABBITMQ_PRODUCER_KEY")

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := rabbitmq.QueueDeclare(ch)
	if err != nil {
		log.Fatal(err)
	}

	ch.QueueBind(q.Name, key, exName, false, nil)

	rabbitmq.Publish(ch, "Deu certo irmão", exName, "text/plain", key)
}
