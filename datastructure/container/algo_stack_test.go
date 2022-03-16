package container

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(1)
	stack.Push(2)

	stack.Push(2)
	stack.Push(2)
	stack.PushList([]Element{2, 4, 5})
	stack.Show()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println("kkkk")
}
