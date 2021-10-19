package main

import (
	"prodconsum-go/consumer"
	"prodconsum-go/producer"
)

func main() {
	// producer
	p := producer.NewProducer{}
	p.CreateProducer()

	// consumer
	c := consumer.NewConsumer{}
	c.CreateConsumer()
}
