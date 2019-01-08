package utils

import (
	"errors"
	"fmt"
)

type Node struct {
	// TODO 后面支持更多数据结构
	data map[string]string
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewList() *LinkedList {
	list := &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}

	return list
}

func (list *LinkedList) append(node *Node) error {
	if node == nil {
		return errors.New("empty node")
	}

	node.next = nil
	if list.size == 0 {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.size++
	list.tail = node

	return nil
}

func (list *LinkedList) get(key string) (string, error) {
	n := list.size
	if n == 0 {
		return "", errors.New("linked list is empty")
	}

	curNode := list.head
	for i := 0; i < n; i++ {
		if val, ok := curNode.data[key]; ok {
			return val, nil
		}
		curNode = curNode.next
	}

	info := fmt.Sprintf("not found %s", key)
	return "", errors.New(info)
}

func (list *LinkedList) del(key string) error {

	return nil
}
