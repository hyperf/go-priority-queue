package go_priority_queue

import (
	"container/heap"
)

type PriorityElement[T any] struct {
	Priority []int64
	Element  T
}

type PriorityQueue[T any] struct {
	Elements []*PriorityElement[T]
}

func (q *PriorityQueue[T]) Insert(element T, priority int64) bool {
	score := []int64{priority}

	q.Elements = append(q.Elements, &PriorityElement[T]{Priority: score, Element: element})

	return true
}

func (q *PriorityQueue[T]) Inserts(element T, priority []int64) bool {
	q.Elements = append(q.Elements, &PriorityElement[T]{Priority: priority, Element: element})

	return true
}

func (q *PriorityQueue[T]) Sort() {
	heap.Init(q)
	var element []*PriorityElement[T]
	for q.Len() > 0 {
		element = append(element, heap.Pop(q).(*PriorityElement[T]))
	}
	q.Elements = element
}

func (q *PriorityQueue[T]) ToArray() []T {
	q.Sort()
	elements := q.Elements
	var result []T
	for _, element := range elements {
		result = append(result, element.Element)
	}
	return result
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.Elements)
}

func (q *PriorityQueue[T]) Less(i, j int) bool {
	return !lessThan(q.Elements[i].Priority, q.Elements[j].Priority)
}

func (q *PriorityQueue[T]) Swap(i, j int) {
	q.Elements[i], q.Elements[j] = q.Elements[j], q.Elements[i]
}

func (q *PriorityQueue[T]) Push(x any) {
	q.Elements = append(q.Elements, x.(*PriorityElement[T]))
}

func (q *PriorityQueue[T]) Pop() any {
	l := len(q.Elements)
	x := q.Elements[l-1]
	q.Elements = q.Elements[0 : l-1]
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
