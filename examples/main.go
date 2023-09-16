package main

import (
	"fmt"
	q "github.com/hyperf/go-priority-queue"
)

func main() {
	queue := &q.PriorityQueue[string]{}
	queue.Insert("Hyperf", 1)
	queue.Insert("Swoole", 2)
	queue.Insert("Swow", 0)

	fmt.Print(queue.ToArray())
}
