package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func Receive() {
	// create kafka config
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// make a connection to kafka server
	connection, err := sarama.NewConsumer([]string{KafkaServerAddress}, config)

	if err != nil {
		log.Fatalf("Fatal error at creating kafka consumer : %v", err)
	}

	// consume from topic
	consumer, consumerErr := connection.ConsumePartition(KafkaTopic, 0, sarama.OffsetNewest)

	if consumerErr != nil {
		log.Fatalf("Fatal at consume message from kafka : %v", consumerErr)
	}

	for message := range consumer.Messages() {
		log.Printf("Receiving message >>> key : %s, value : %s", message.Key, message.Value)
	}

}
