package utils

import (
	"errors"
	"fmt"
)

type Node struct {
	// TODO 后面支持更多数据结构
	Data map[string]string
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewList() *LinkedList {
	list := &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}

	return list
}

func (list *LinkedList) append(node *Node) error {
	// TODO 是否需要保证key的唯一性
	if node == nil {
		return errors.New("empty node")
	}

	node.Next = nil
	if list.Size == 0 {
		list.Head = node
	} else {
		list.Tail.Next = node
	}
	list.Size++
	list.Tail = node

	return nil
}

func (list *LinkedList) prefix(node *Node) error {
	if node == nil {
		return errors.New("empty node")
	}

	if list.Size == 0 {
		list.Tail = node
	}

	node.Next = list.Head
	list.Head = node
	list.Size++

	return nil
}

func (list *LinkedList) exchange(nodeA, nodeB *Node) error {
	// TODO
	return nil
}

func (list *LinkedList) setFront(preNode, node *Node) error {
	n := list.Size
	if n == 0 {
		return errors.New("linked list is empty")
	}

	if list.Head == node {
		return nil
	}
	if list.Tail == node {
		list.Tail = preNode
	}

	//fmt.Println(">>>check next", preNode, node)
	preNode.Next = node.Next
	node.Next = list.Head
	list.Head = node

	return nil
}

func (list *LinkedList) SetFront(key string) error {
	n := list.Size
	if n == 1 {
		return nil
	}

	preNode := list.Head
	curNode := list.Head.Next
	if _, ok := preNode.Data[key]; ok {
		return nil
	}
	for i := 1; i < n; i++ {
		if _, ok := curNode.Data[key]; ok {
			err := list.setFront(preNode, curNode)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		preNode = curNode
		curNode = curNode.Next
	}

	return errors.New(fmt.Sprintf("not found %s", key))
}

func (list *LinkedList) SetTail(key, val string) error {
	node := &Node{
		Data: map[string]string{key: val},
		Next: nil,
	}

	err := list.append(node)
	if err != nil {
		return err
	}

	return nil
}

func (list *LinkedList) SetHead(key, val string) error {
	node := &Node{
		Data: map[string]string{key: val},
	}

	if err := list.prefix(node); err != nil {
		return err
	}

	return nil
}

func (list *LinkedList) Get(key string) (string, error) {
	n := list.Size
	if n == 0 {
		return "", errors.New("linked list is empty")
	}

	curNode := list.Head
	for i := 0; i < n; i++ {
		if val, ok := curNode.Data[key]; ok {
			return val, nil
		}
		curNode = curNode.Next
	}

	info := fmt.Sprintf("not found %s", key)
	return "", errors.New(info)
}

func (list *LinkedList) Del(key string) error {
	n := list.Size
	if n == 0 {
		return errors.New("linked list is empty")
	}

	// 判断head是否是目标
	if _, ok := list.Head.Data[key]; ok {
		//tmp := list.head.next
		//list.head.next = nil
		//list.head = tmp

		// TODO 这种写法有待确认
		list.Head, list.Head.Next = list.Head.Next, nil
		list.Size--

		return nil
	}

	preNode := list.Head
	curNode := list.Head.Next
	for i := 1; i < n; i++ {
		if key != curNode.Data[key] {
			preNode = curNode
			curNode = curNode.Next
			continue
		}

		preNode.Next = curNode.Next
		curNode.Next = nil
		list.Size--
	}

	return nil
}

func (list *LinkedList) DelHead() error {
	n := list.Size
	if n == 0 {
		return errors.New("linked list is empty")
	}

	list.Head, list.Head.Next = list.Head.Next, nil
	list.Size--
	return nil
}

func (list *LinkedList) DelTail() error {
	n := list.Size
	if n == 0 {
		return errors.New("linked list is empty")
	}

	if n == 1 {
		list.Head = nil
		return nil
	}

	preNode := list.Head
	curNode := list.Head.Next
	for i := 1; i < n; i++ {
		preNode = curNode
		curNode = curNode.Next
	}

	list.Tail = preNode
	preNode.Next = nil
	list.Size--

	return nil
}
