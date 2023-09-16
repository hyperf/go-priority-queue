package main

import (
	"fmt"
	"github.com/hyperf/go-priority-queue"
)

func main() {
	queue := &go_priority_queue.PriorityQueue[string]{}
	queue.Insert("Hyperf", 1)
	queue.Insert("Swoole", 2)
	queue.Insert("Swow", 0)

	fmt.Print(queue.ToArray())
}
