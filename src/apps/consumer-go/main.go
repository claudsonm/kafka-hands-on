package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
		"group.id":          "go-consumer-app",
		"auto.offset.reset": "earliest",
		"broker.address.family": "v4",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topico-exemplo", "^aRegex.*[Tt]opico"}, nil)

	fmt.Println("Aguardando por novas mensagens...")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Mensagem lida de %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Erro ao consumir: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
