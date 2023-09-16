package go_priority_queue

import (
	"container/heap"
)

type PriorityElement[T any] struct {
	Priority []int64
	Element  T
}

type PriorityQueue[T any] struct {
	elements []*PriorityElement[T]
	priority int64
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (q *PriorityQueue[T]) Insert(element T, priority int64) bool {
	q.priority--

	score := []int64{priority, q.priority}

	q.elements = append(q.elements, &PriorityElement[T]{Priority: score, Element: element})

	return true
}

func (q *PriorityQueue[T]) Inserts(element T, priority []int64) bool {
	q.priority--
	priority = append(priority, q.priority)
	q.elements = append(q.elements, &PriorityElement[T]{Priority: priority, Element: element})

	return true
}

func (q *PriorityQueue[T]) Sort() {
	heap.Init(q)
	var element []*PriorityElement[T]
	for q.Len() > 0 {
		element = append(element, heap.Pop(q).(*PriorityElement[T]))
	}
	q.elements = element
}

func (q *PriorityQueue[T]) ToArray() []T {
	q.Sort()
	elements := q.elements
	var result []T
	for _, element := range elements {
		result = append(result, element.Element)
	}
	return result
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.elements)
}

func (q *PriorityQueue[T]) Less(i, j int) bool {
	return !lessThan(q.elements[i].Priority, q.elements[j].Priority)
}

func (q *PriorityQueue[T]) Swap(i, j int) {
	q.elements[i], q.elements[j] = q.elements[j], q.elements[i]
}

func (q *PriorityQueue[T]) Push(x any) {
	q.elements = append(q.elements, x.(*PriorityElement[T]))
}

func (q *PriorityQueue[T]) Pop() any {
	l := len(q.elements)
	x := q.elements[l-1]
	q.elements = q.elements[0 : l-1]
	return x
}

func lessThan(s1 []int64, s2 []int64) bool {
	l1 := len(s1)
	l2 := len(s2)

	l := min(l1, l2)
	for i := 0; i < l; i++ {
		if s1[i] == s2[i] {
			continue
		}

		return s1[i] < s2[i]
	}

	return l1 < l2
}
