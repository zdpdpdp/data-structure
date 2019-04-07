package queue

import "errors"

var emptyError = errors.New("stack is empty")

type Queue struct {
	container []int
}

func NewQueue() *Queue {
	return &Queue{
		container: make([]int, 0),
	}
}

func (q *Queue) Push(data int) {
	q.container = append(q.container, data)
}

func (q *Queue) Pop() (int, error) {
	if len(q.container) == 0 {
		return 0, emptyError
	}

	res := q.container[0]

	q.container = q.container[1:]

	return res, nil
}

func (q *Queue) len() int {
	return len(q.container)
}
