package linklist

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
* 打印单链表
 */
func printListNode(node *ListNode) {
	p := node
	for p != nil {
		fmt.Print(strconv.Itoa(p.Val) + " ")
		p = p.Next
		if p != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func printNode(node *ListNode) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	fmt.Println(strconv.Itoa(node.Val))
}
