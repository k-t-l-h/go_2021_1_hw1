package internal

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
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

func Calc(in io.Reader) (int, error) {
	numbers := NewQueue()
	operands := NewQueue()

	var (
		input   string
		br      bool
		builder strings.Builder
	)

	for {
		_, err := fmt.Fscan(in, &input)
		if err != nil {
			if err == io.EOF && numbers.Len() == 0 {
				break
			}
			panic(err)
		}

		if br && input != ")" {
			builder.WriteString(input)
			builder.WriteString(" ")
		} else {
			switch input {
			case "+":
				operands.Push(input)
			case "-":
				operands.Push(input)
			case "*":
				operands.Push(input)
			case "/":
				operands.Push(input)
			case "=":
				if numbers.Len() != 1 && operands.Len() == 0 {
					return 0, errors.New("multiple answer")
				}
				for operands.Len() != 0 {
					switch operands.Pop() {
					case "+":
						number := numbers.Pop().(int) + numbers.Pop().(int)
						numbers.Push(number)
					case "-":
						number := -numbers.Pop().(int) + numbers.Pop().(int)
						numbers.Push(number)
					case "*":
						number := numbers.Pop().(int) * numbers.Pop().(int)
						numbers.Push(number)
					case "/":
						down := numbers.Pop().(int)
						numbers.Push(numbers.Pop().(int) / down)
					}
				}
				return numbers.Pop().(int), nil
			case "(":
				br = true
			case ")":
				if !br {
					return 0, errors.New("brackets error")
				}
				br = !br
				builder.WriteString("= ")
				subres, _ := Calc(strings.NewReader(builder.String()))
				builder.Reset()
				numbers.Push(subres)
				input = strconv.Itoa(subres)
			default:
				n, err := strconv.Atoi(input)
				if err != nil {
					return 0, errors.New("wrong type provided")
				}
				numbers.Push(n)
				if numbers.Len()%2 == 0 {
					switch operands.Pop() {
					case "+":
						number := numbers.Pop().(int) + numbers.Pop().(int)
						numbers.Push(number)
					case "-":
						number := -numbers.Pop().(int) + numbers.Pop().(int)
						numbers.Push(number)
					case "*":
						number := numbers.Pop().(int) * numbers.Pop().(int)
						numbers.Push(number)
					case "/":
						down := numbers.Pop().(int)
						numbers.Push(numbers.Pop().(int) / down)
					}
				}
			}
		}
	}
	return 0, nil
}
