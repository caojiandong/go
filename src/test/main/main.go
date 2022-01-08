package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		fmt.Println(num)
		sum += num
	}

	// 可以这么理解，range会返回两个值，一个是key，一个是value
	// array也是一种特殊的map。
	// go没有arr，只有map.
	// array[1, 2, 3, 4]
	// 可以表达为map["1":1, "2":2, "3":3, "4":4]
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
