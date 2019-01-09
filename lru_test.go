package midcache

import (
	"fmt"
	"testing"
)

func TestLRU_Set(t *testing.T) {
	lru := NewLRU(3)

	lru.Set("nihao", "shijie")
	lru.Set("hello", "world")
	lru.Set("bonjour", "monde")
	lru.Set("say", "hello")

	fmt.Println(lru.list.Head, lru.list.Tail)
	val := lru.Get("nihao")

	fmt.Println("bonjour: ", val)

	fmt.Println(lru.list.Head, lru.list.Tail)
}

func TestLRU_Get(t *testing.T) {

}
