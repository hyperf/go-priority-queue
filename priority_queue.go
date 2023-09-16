package go_priority_queue

type PriorityElement[T any] struct {
	Priority []int64
	Element  T
}

type PriorityQueue[T any] struct {
	Elements []*PriorityElement[T]
}

func (q *PriorityQueue[T]) Insert(element any, priority int64) bool {
	score := []int64{priority}

	q.Elements = append(q.Elements, &PriorityElement[T]{Priority: score, Element: element})

	return true
}

func (q *PriorityQueue[T]) Inserts(element any, priority []int64) bool {
	q.Elements = append(q.Elements, &PriorityElement[T]{Priority: priority, Element: element})

	return true
}

func (q PriorityQueue[T]) toSlice() {

}

func (q *PriorityQueue[T]) Len() int {
	return len(q.Elements)
}

func (q *PriorityQueue[T]) Less(i, j int) bool {
	return lessThan(q.Elements[i].Priority, q.Elements[j].Priority)
}

func (q *PriorityQueue[T]) Swap(i, j int) {
	q.Elements[i], q.Elements[j] = q.Elements[j], q.Elements[i]
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
