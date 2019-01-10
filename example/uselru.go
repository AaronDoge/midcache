package main

import (
	"fmt"
	"midcache"
	"time"
)

const TEST_TIMES = 1000000

func main()  {

	cache := midcache.NewLRU(200)
	startSet := time.Now()
	for i := 0; i < TEST_TIMES; i++ {
		key := fmt.Sprintf("No.%d", i % 200)
		val := fmt.Sprintf("values of %s", key)
		//fmt.Println(">>>check key", key, "val", val)
		err := cache.SetNoRe(key, val)
		if err != nil {
			fmt.Println("set error", err.Error(), ", location", key)
		}
	}
	endSet := time.Now()
	//fmt.Println("Set 总用时：", endSet.Sub(startSet))
	//
	//fmt.Println("Head", cache.List.Head, "tail", cache.List.Tail)

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

	fmt.Println("set总用时：", endSet.Sub(startSet))
	fmt.Println("get总用时：", endTime.Sub(startTime))
	// 100w 	11s
	// 10w 		1.15s
}




