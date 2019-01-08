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
	val := lru.Get("hello")

	fmt.Println("bonjour: ", val)


}

func TestLRU_Get(t *testing.T) {

}
