package kafka

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

/*
 *	Kafka Server Config
 */
const (
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

func Send(busNo int, message string) {
	// create kafka server config
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// make a producer
	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)

	if err != nil {
		log.Fatalf("Kafka Producer Error >>> %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(busNo)),
		Value: sarama.StringEncoder(message),
	}

	// Drop message to kafka topic
	producer.SendMessage(msg)
}
