package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	const d = 500000
	const n = 3e20/d

	fmt.Println(n/d)
	fmt.Println(int64(n))
	fmt.Println(math.Sin(d))
}