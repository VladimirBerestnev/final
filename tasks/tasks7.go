package tasks

import (
	"errors"
	"fmt"
)

//Реализовать структуру данных &quot;стек&quot; с функциональностью Push, Pop. Помещение элемента в стек добавляет его в самую верхнюю позицию, а удаление из стека извлекает самый верхний элемент. Top возвращает самый верхний элемент в стеке. Если стек пуст, он возвращает нулевое значение и ошибку, говорящую о том, что стек пуст. IsEmpty возвращает true, если стек пуст, и false в противном случае. Print итерируется по стеку и выводит элементы

type Stack struct {
	stack []int
}

func NewStack(s []int) *Stack {
	return &Stack{stack: s}
}

func (s *Stack) Push(el int) {
	s.stack = append(s.stack, el)
}

func (s *Stack) Pop() {
	if len(s.stack) == 0 {
		fmt.Println("Stack is Empty")
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) Top() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("Stack is Empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Print() {
	for _, el := range s.stack {
		fmt.Print(el)
	}
}
