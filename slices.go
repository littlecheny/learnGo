package main

import(
	"fmt"
)

func main() {
	twoD := make([][]int, 3)
	for i := range 3{
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i+ j
		}
	}
	fmt.Println(twoD)
}