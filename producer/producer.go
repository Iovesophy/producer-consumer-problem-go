package producer

import "fmt"

var (
	_ Producer = (*NewProducer)(nil)
)

type NewProducer struct{}

type Producer interface {
	CreateProducer()
}

func (p *NewProducer) CreateProducer() {
	fmt.Println("producer")
}
