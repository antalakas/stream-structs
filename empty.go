package stream_structs

import "github.com/confluentinc/confluent-kafka-go/kafka"

type EmptyInterface interface {
}

type KafkaChannelWrapper struct {
	KafkaChannel chan *kafka.Message
	AssignedPartitionsChannel chan string
}
