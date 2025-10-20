package main

import (
	"fmt"
)

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				for k := i; k > j; k-- {
					arr[k], arr[k-1] = arr[k-1], arr[k]
				}
			}
		}
	}
}

func merger(a, b []int) {
	for _, v := range b {
		a = append(a, v)
	}
}

func main() {
	//输入一个数组
	var arr []int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	if len(arr) == 0 {
		fmt.Println("数组为空")
		return
	}
	sort(arr)
	fmt.Println("排序后的数组为：", arr)
	merger(arr, []int{1, 2, 3})
	fmt.Println("合并后的数组为：", arr)
}
