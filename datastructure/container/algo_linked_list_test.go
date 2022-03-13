package container

import (
	"testing"
)

func TestNewLinklist(t *testing.T) {
	linkedlist := NewLinklist()
	linkedlist.AddAtHead(1)
	linkedlist.AddAtTail(2)
	linkedlist.AddAtHead(3)
	linkedlist.AddAtIndex(1, 4)
	linkedlist.DeleteAtIndex(1)
	linkedlist.DeleteAtIndex(4)
	linkedlist.DeleteAtIndex(2)
	linkedlist.AddAtIndex(-1, 22)
	linkedlist.AddAtTail(33)
	linkedlist.AddAtHead(99)
	linkedlist.Show()
}
