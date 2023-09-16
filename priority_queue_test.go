package priority_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue_ToArray(t *testing.T) {
	queue := &PriorityQueue[string]{}
	queue.Insert("Hyperf", 1)
	queue.Insert("Swoole", 2)
	queue.Insert("Swow", 0)

	assert.Equal(t, []string{"Swoole", "Hyperf", "Swow"}, queue.ToArray())
	assert.Equal(t, []string{"Swoole", "Hyperf", "Swow"}, queue.ToArray())
}

func TestPriorityQueue_SamePriority(t *testing.T) {
	queue := &PriorityQueue[string]{}
	queue.Insert("Hyperf", 1)
	queue.Insert("Swoole", 2)
	queue.Insert("Swow", 2)

	assert.Equal(t, []string{"Swoole", "Swow", "Hyperf"}, queue.ToArray())
}
