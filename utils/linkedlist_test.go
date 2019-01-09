package utils

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()

	node1 := &Node{
		data: map[string]string{"nihao": "shijie"},
		next: nil,
	}
	node2 := &Node{
		data: map[string]string{"hello": "world"},
	}

	l.append(node1)
	l.append(node2)

	val, _ := l.Get("nihao")
	val2, _ := l.Get("hello")

	fmt.Println("nihao: ", val)
	fmt.Println("hello: ", val2)
}
