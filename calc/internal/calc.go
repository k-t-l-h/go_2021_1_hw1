package internal

import (
	"io"
)

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{data: make([]interface{}, 0)}
}

func (q *Queue) Push(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Pop() interface{} {
	if len(q.data) == 0 {
		panic("no data found")
	}
	value := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return value
}

func (q *Queue) Len() int {
	return len(q.data)
}

func Calc(reader io.Reader) (int,error) {
	return 0, nil
}
