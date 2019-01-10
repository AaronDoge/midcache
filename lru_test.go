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

	fmt.Println("size", lru.List.Size)
	fmt.Println("head:", lru.List.Head, ", tail:", lru.List.Tail)
	val := lru.Get("nihao")

	fmt.Println("say: ", val)

	fmt.Println(lru.List.Head, lru.List.Tail)
	fmt.Println("size", lru.List.Size)
}

func TestLRU_Get(t *testing.T) {

}
