package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	hoge []string
	mu   sync.Mutex
)

func initialize() {
	rand.Seed(time.Now().UnixNano())
}

func item() string {
	list := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return string(list[rand.Intn(len(list))])
}

func p() {
	mu.Lock()
	if len(hoge) == 0 {
		v := item()
		hoge = append(hoge, v)
		fmt.Println(hoge, "<", v)
	}
	mu.Unlock()
}

func c() {
	mu.Lock()
	if len(hoge) >= 1 {
		v := hoge[len(hoge)-1]
		hoge = hoge[:len(hoge)-1]
		fmt.Println(hoge, ">", v)
	}
	mu.Unlock()
}

func main() {
	initialize()
	for {
		go c()
		go p()
	}
}
