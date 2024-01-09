package kafka

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

const (
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

func Send(busNo int, message string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)

	if err != nil {
		log.Fatalf("Kafka Producer Error >>> %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(busNo)),
		Value: sarama.StringEncoder(message),
	}

	producer.SendMessage(msg)
}
