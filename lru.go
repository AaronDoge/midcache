package midcache

import (
	"fmt"
	"midcache/utils"
)

type LRU struct {
	list 	*utils.LinkedList
	maxSize int
}

func NewLRU(size int) *LRU {
	li := utils.NewList()
	lru := &LRU{
		list:	li,
		maxSize: size,
	}

	return lru
}

func (lru *LRU) Set(key, val string) error {
	// 先写再删
	err := lru.list.SetHead(key, val)
	if err != nil {
		return err
	}

	if lru.list.Size > lru.maxSize {
		err  := lru.list.DelTail()
		if err != nil {
			return err
		}
	}
	return nil
}

func (lru *LRU) Get(key string) string {
	val, err := lru.list.Get(key)
	if err != nil {
		fmt.Println("get val error. key:", key, "error: ", err.Error())
		return ""
	}

	err = lru.list.SetFront(key)
	if err != nil {
		fmt.Println("get val error. key:", key, "error: ", err.Error())
		return ""
	}

	return val
}


