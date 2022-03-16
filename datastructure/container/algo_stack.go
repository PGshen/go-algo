package container

import (
	"fmt"
)

// 定义元素，可保存任意类型的值
type Element interface{}

type Stack struct {
	list []Element
}

func NewStack() *Stack {
	return &Stack{
		list: make([]Element, 0),
	}
}

func (s *Stack) Len() int {
	return len(s.list)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(x interface{}) {
	s.list = append(s.list, x)
}

func (s *Stack) PushList(x []Element) {
	s.list = append(s.list, x...)
}

func (s *Stack) Pop() Element {
	if s.IsEmpty() {
		fmt.Println("Stack is empty!")
		return nil
	}
	ret := s.list[s.Len()-1]
	s.list = s.list[:s.Len()-1]
	return ret
}

func (s *Stack) Top() Element {
	if s.IsEmpty() {
		fmt.Println("Stack is empty!")
		return nil
	}
	return s.list[s.Len()-1]
}

func (s *Stack) Clear() bool {
	if s.IsEmpty() {
		return true
	}
	for i := 0; i < s.Len(); i++ {
		s.list[i] = nil
	}
	s.list = make([]Element, 0)
	return true
}

func (s *Stack) Show() {
	len := s.Len()
	for i := len - 1; i >= 0; i-- {
		fmt.Print(s.list[i])
		fmt.Print(" ")
	}
	fmt.Println()
}
