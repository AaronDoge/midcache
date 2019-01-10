package midcache

import (
	"fmt"
	"midcache/utils"
)

type LRU struct {
	List    *utils.LinkedList
	maxSize int
}

func NewLRU(size int) *LRU {
	li := utils.NewList()
	lru := &LRU{
		List:    li,
		maxSize: size,
	}

	return lru
}

func (lru *LRU) Set(key, val string) error {
	// 先写再删
	err := lru.List.SetHead(key, val)
	if err != nil {
		return err
	}

	if lru.List.Size > lru.maxSize {
		err := lru.List.DelTail()
		if err != nil {
			return err
		}
	}
	return nil
}

func (lru *LRU) SetNoRe(key, val string) error {
	if lru.List.Size != 0 {
		_, node, err := lru.List.GetNode(key)
		if err == nil && node != nil {
			node.Data[key] = val
			return nil
		}
	}

	err := lru.Set(key, val)
	if err != nil {
		return err
	}
	return nil
}

func (lru *LRU) Get(key string) string {
	pre, node, err := lru.List.GetNode(key)
	if err != nil {
		fmt.Println("get val error. key:", key, "error: ", err.Error())
		return ""
	}

	err = lru.List.SetFrontNode(pre, node)
	//err = lru.List.SetFront(key)
	if err != nil {
		fmt.Println("set front error. key:", key, "error: ", err.Error())
		return ""
	}

	return node.Data[key]
}
