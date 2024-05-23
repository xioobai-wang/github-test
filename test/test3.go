package main

import (
	"fmt"
	"sort"
)

// 结构体定义
type test struct {
	value int
	str   string
}

func main() {
	s := make([]test, 5)
	s[0] = test{value: 2, str: "test2"}
	s[1] = test{value: 4, str: "test4"}
	s[2] = test{value: 1, str: "test1"}
	s[3] = test{value: 5, str: "test5"}
	s[4] = test{value: 3, str: "test3"}
	fmt.Println("初始化结果:")
	fmt.Println(s)

	// 从小到大排序(稳定排序)
	sort.SliceStable(s, func(i, j int) bool {
		if s[i].value < s[j].value {
			return true
		}
		return false
	})
	fmt.Println(s)
}

