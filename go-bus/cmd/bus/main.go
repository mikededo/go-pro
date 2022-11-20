package main

import (
	"fmt"
	"math/rand"
	"time"

	bus "github.com/mddg/go-bus/internal"
)

func genHandler(id int) func(...int) {
	return func(e ...int) {
		fmt.Printf("[handler %d] received %v\n", id, e)
	}
}

func main() {
	b := bus.NewEventBus(1)
	h1 := genHandler(1)
	h2 := genHandler(2)
	h3 := genHandler(3)

	b.Subscribe("num", h1)
	b.Subscribe("num", h2)
	b.Subscribe("num", h3)

	b.Subscribe(fmt.Sprintf("num_%d", 1), h1)
	b.Subscribe(fmt.Sprintf("num_%d", 2), h2)
	b.Subscribe(fmt.Sprintf("num_%d", 3), h3)

	b.Subscribe("multinum", h2)

	b.Notify("num", 12)
	b.Notify("multinum", 1, 2, 3, 4)

	for i := 0; i <= 11; i++ {
		b.Notify(fmt.Sprintf("num_%d", rand.Intn(3)+1), i)
		time.Sleep(500 * time.Millisecond)
	}

	b.Clear("num")
	b.Clear("multinum")
}
