package main

import (
	"fmt"

	"github.com/andrzejewsky/go-queue/queues"
)

func main() {

	queue, _ := queues.GetQueue("redis", map[string]string{
		"host": "192.168.99.100",
		"port": "6379",
	})

	queue.Push("test")
	value, _ := queue.Pop()
	fmt.Println(value)
	defer queue.Close()

}
