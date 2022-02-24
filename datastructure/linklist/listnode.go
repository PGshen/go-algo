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
func printNode(node *ListNode) {
	p := node
	for p != nil {
		fmt.Print(strconv.Itoa(p.Val) + " ")
		p = p.Next
	}
	fmt.Println()
}
