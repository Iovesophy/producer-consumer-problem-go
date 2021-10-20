package consumer

import "fmt"

var (
	_ Consumer = (*NewSampleConsumer)(nil)
)

type NewSampleConsumer struct {
	Msgs  *chan int
	Count int
}

type Consumer interface {
	CreateSampleConsumer()
}

// Set Visibility-Flow of Consumer recieving chan value for test
func (c *NewSampleConsumer) CreateSampleConsumer() {
	fmt.Println("start consumer")
	for {
		msg := <-*c.Msgs
		fmt.Println("consume: Received:", msg)
		c.Count++
	}
}
