package midcache

import (
	"fmt"
	"testing"
)

func TestLRU_Set(t *testing.T) {
	lru := NewLRU(4)

	lru.Set("nihao", "shijie")
	lru.Set("hello", "world")
	lru.Set("bonjour", "monde")
	lru.Set("say", "hello")

	fmt.Println("size", lru.list.Size)
	fmt.Println("head:", lru.list.Head, ", tail:", lru.list.Tail)
	val := lru.Get("nihao")

	fmt.Println("say: ", val)

	fmt.Println(lru.list.Head, lru.list.Tail)
	fmt.Println("size", lru.list.Size)
}

func TestLRU_Get(t *testing.T) {

}
