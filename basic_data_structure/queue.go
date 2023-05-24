package basic_data_structure

type Queue struct {
	queue []int
}

func (q *Queue) Enqueue(value int) {
	q.queue = append(q.queue, value)
}

func (q *Queue) Dequeue() int {
	value := q.queue[0]
	q.queue = q.queue[1:]

	return value
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
