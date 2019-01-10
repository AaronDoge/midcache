package main

import (
	"fmt"
	"midcache"
	"time"
)

const TEST_TIMES = 100000

func main()  {

	cache := midcache.NewLRU(200)

	for i := 0; i < 200; i++ {
		key := fmt.Sprintf("No.%d", i)
		val := fmt.Sprintf("values of %s", key)
		cache.Set(key, val)
	}

	//fmt.Println("Size:", cache.List.Size)
	//fmt.Println("head", cache.List.Head, "tail", cache.List.Tail)
	//
	//fmt.Println("get No.321: ", cache.Get("No.321"))
	//fmt.Println("head", cache.List.Head, "tail", cache.List.Tail)
	//
	//fmt.Println("get No.300: ", cache.Get("No.300"))
	//fmt.Println("head", cache.List.Head, "tail", cache.List.Tail)
	//
	//fmt.Println("get No.299: ", cache.Get("No.299"))
	//fmt.Println("head", cache.List.Head, "tail", cache.List.Tail)


	startTime := time.Now()
	for i := 0; i < TEST_TIMES; i++ {
		getKey := fmt.Sprintf("No.%d", i % 200)
		fmt.Println("获取结果：", cache.Get(getKey))
	}
	endTime := time.Now()

	fmt.Println("总用时：", endTime.Sub(startTime))
	// 100w 	11s
	// 10w 		1.15s
}




