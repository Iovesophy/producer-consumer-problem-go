package producer

import (
	"fmt"
)

var (
	_ Producer = (*NewSampleProducer)(nil)
)

// Producer interface list
type Producer interface {
	CreateSampleProducer(int)
}

// Producer Prams of Value
type NewSampleProducer struct {
	Msgs *chan int
	Done *chan bool
}

// Set Visibility-Flow of Producer sending chan value for test
func (p *NewSampleProducer) CreateSampleProducer(max int) {
	fmt.Println("start producer")
	for ch := 0; ch < max; ch++ {
		fmt.Println("produce: Sending ", ch)
		*p.Msgs <- ch
	}
	*p.Done <- true
}
