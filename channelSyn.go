package main

import(
	"fmt"
	"time"
)

func work(done chan bool){
	fmt.Println("Start...")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true
}

func main() {
	done := make(chan bool)
	go work(done)
	<- done
}