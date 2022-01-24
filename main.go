package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	buf []string
	mu  sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func item() string {
	return string(rand.Intn(26) + 'A')
}

func producer() {
	defer mu.Unlock()
	mu.Lock()
	if len(buf) == 0 {
		v := item()
		buf = append(buf, v)
		fmt.Println(buf, "<", v)
	}
}

func consumer() {
	defer mu.Unlock()
	mu.Lock()
	if len(buf) >= 1 {
		v := buf[len(buf)-1]
		buf = buf[:len(buf)-1]
		fmt.Println(buf, ">", v)
	}
}

func main() {
	for {
		producer()
		consumer()
	}
}
