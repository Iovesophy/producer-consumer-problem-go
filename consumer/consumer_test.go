package consumer

import (
	"prodconsum-go/producer"
)

func ExampleConsumer_callSampleConsumerInterfaceTest() {
	// set flag value of testdata
	max := 10

	// make region
	msgs := make(chan int)
	done := make(chan bool)

	// consumer
	c := NewSampleConsumer{}
	c.Msgs = &msgs
	go c.CreateSampleConsumer()

	// producer
	p := producer.NewSampleProducer{}
	p.Msgs = &msgs
	p.Done = &done
	go p.CreateSampleProducer(max)
	for {
		if c.Count >= max {
			<-*p.Done
			break
		}
	}
	// Unordered Output:
	// start consumer
	// consume: Received: 0
	// consume: Received: 1
	// consume: Received: 2
	// consume: Received: 3
	// consume: Received: 4
	// consume: Received: 5
	// consume: Received: 6
	// consume: Received: 7
	// consume: Received: 8
	// consume: Received: 9
	// start producer
	// produce: Sending  0
	// produce: Sending  1
	// produce: Sending  2
	// produce: Sending  3
	// produce: Sending  4
	// produce: Sending  5
	// produce: Sending  6
	// produce: Sending  7
	// produce: Sending  8
	// produce: Sending  9
}
