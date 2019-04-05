package trees

import "errors"

var emptyError = errors.New("stack is empty")

type Queue struct {
	container []*Node
}

func NewQueue() *Queue {
	return &Queue{
		container: make([]*Node, 0),
	}
}

func (q *Queue) Push(data *Node) {
	q.container = append(q.container, data)
}

func (q *Queue) Pop() (*Node, error) {
	if len(q.container) == 0 {
		return nil, emptyError
	}

	res := q.container[0]

	q.container = q.container[1:]

	return res, nil
}

func (q *Queue) len() int {
	return len(q.container)
}
