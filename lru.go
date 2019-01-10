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

func (lru *LRU) Get(key string) string {
	val, err := lru.List.Get(key)
	if err != nil {
		fmt.Println("get val error. key:", key, "error: ", err.Error())
		return ""
	}

	// TODO List.Get()中一次遍历，这里一次遍历，有点浪费，可以直接利用Get中返回的结果；这需要List.Get返回一个*Node类型了
	err = lru.List.SetFront(key)
	if err != nil {
		fmt.Println("set front error. key:", key, "error: ", err.Error())
		return ""
	}

	return val
}
