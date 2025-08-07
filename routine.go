package main

import(
	"fmt"
	"time"
)

func f(form string){
	i := 0
	for i<3 {
		fmt.Println(form, ":", "i")
		time.Sleep(100*time.Millisecond)
		i++
	}
}

func main() {
	f("direct")

	go f("goroutine")

	time.Sleep(100*time.Millisecond)

	go func (msg string){
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
}