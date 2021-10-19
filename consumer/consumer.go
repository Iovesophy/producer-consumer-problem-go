package consumer

import "fmt"

var (
	_ Consumer = (*NewConsumer)(nil)
)

type NewConsumer struct{}

type Consumer interface {
	CreateConsumer()
}

func (c *NewConsumer) CreateConsumer() {
	fmt.Println("consumer")
}
