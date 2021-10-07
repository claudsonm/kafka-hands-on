package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "go-consumer-app",
		"auto.offset.reset": "earliest",
		"broker.address.family": "v4",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opico"}, nil)
	// c.SubscribeTopics([]string{"myTopic"}, nil)

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
