package main

import (
	"fmt"
	"midcache"
)

func main() {
	cache := midcache.NewLRU(200)

	for i := 0; i < 200; i++ {
		key := fmt.Sprintf("No.%d", i)
		val := fmt.Sprintf("values of %s", key)
		cache.Set(key, val)
	}

	cache.Set("No.123", "helloworld")

	val := cache.Get("No.123")
	fmt.Println("val is", val)
	fmt.Println("head:", cache.List.Head, "tail:", cache.List.Tail)
}

