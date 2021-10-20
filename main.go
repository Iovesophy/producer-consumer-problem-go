package main

import (
	"flag"
	"prodconsum-go/consumer"
	"prodconsum-go/producer"
)

// flag option list
type flagVal struct {
	max *int // defalut = 5
}

// flag option setting and return flag value
func (f *flagVal) getFlag() int {
	f.max = flag.Int("n", 5, "defines the max num of messages")
	flag.Parse()
	return *f.max
}

// main
func main() {
	// get flag value
	f := flagVal{}
	max := f.getFlag()

	msgs := make(chan int)
	done := make(chan bool)

	// consumer
	c := consumer.NewConsumer{}
	c.Msgs = &msgs
	go c.CreateSampleConsumer()

	// producer
	p := producer.NewProducer{}
	p.Msgs = &msgs
	p.Done = &done
	go p.CreateSampleProducer(max)
	for {
		if c.Count >= max {
			<-*p.Done
			break
		}
	}
}
