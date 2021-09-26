package internal

import (
	"fmt"
	"strings"
	"testing"
)

//проверка очереди
func TestNumPushPlace(t *testing.T) {
	num := NewQueue()
	for i := 0; i < 50; i++ {
		num.Push(i)
		if num.Pop().(int) != i {
			t.Errorf("num push is not working with %d elements", i)
		}
	}
}

func TestNumPop(t *testing.T) {
	num := NewQueue()
	for i := 0; i <= 50; i++ {
		num.Push(i)
	}

	for i := 50; i >= 0; i-- {
		j := num.Pop().(int)
		if j != i {
			t.Errorf("expected: %d got: %d", i, j)
		}
	}
}

func TestNumPopError(t *testing.T) {

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("pop should have panicked")
			}
		}()

		num := NewQueue()
		num.Pop()
	}()
}

func TestStrPushPlace(t *testing.T) {
	op := NewQueue()
	for i := 0; i < 50; i++ {
		op.Push(fmt.Sprint(i))
		if op.Pop().(string) != fmt.Sprint(i) {
			t.Errorf("str push is not working with %d elements", i)
		}
	}
}

func TestStrPopError(t *testing.T) {

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("pop should have panicked")
			}
		}()
		op := NewQueue()
		op.Pop()
	}()
}

//проверка базовых операций
func TestCalc(t *testing.T) {
	var cases = []struct {
		expected string
		input    string
	}{
		{
			//сложение
			input:    "1 + 2 + 3 + 4  =",
			expected: "10",
		},
		{
			//вычитание
			input:    "15 - 2 - 3 - 4 = ",
			expected: "6",
		},
		{
			//отрицательные числа
			input:    "3 - 4 =",
			expected: "-1",
		},
		{
			//отрицательные числа
			input:    "-3 + 4 =",
			expected: "1",
		},
		{
			//деление
			input:    "4 / 2 =",
			expected: "2",
		},
		{
			//целочисленное деление
			input:    "5 / 2 =",
			expected: "2",
		},
		{
			//равенство чисел самим себе
			input:    "3 = ",
			expected: "3",
		},
		{
			//обработка пустых запросов
			input:    "",
			expected: "0",
		},
		{
			//обработка скобок
			input:    "( 1 + 2 ) =",
			expected: "3",
		},
		{
			//обработка нескольких скобок
			input:    "( 1  +  2 ) + ( 2 + 1 ) =",
			expected: "6",
		},
	}

	for _, item := range cases {
		ok, _ := Calc(strings.NewReader(item.input))
		if fmt.Sprint(ok) != item.expected {
			t.Error(item, ok)
		}
	}
}
