package utils

// TODO 后面支持更多数据结构
type Node struct {
	Data map[string]string
	Next *Node
}

type NodeMap struct {
	Data map[string]map[string]string
	Next *NodeMap
}

// DLL, Doubly Linked List
type NodeDLL struct {
	Data map[string]string
	Prev *NodeDLL
	Next *NodeDLL
}

type NodeDLLMap struct {
	Data map[string]map[string]string
	Prev *NodeDLL
	Next *NodeDLL
}