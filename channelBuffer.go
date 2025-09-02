package main

import(
	"fmt"
)

func main(){
	message := make(chan string,2)

	message <- "first"
	message <- "second"

	fmt.Println(<-message)
	fmt.Println(<-message)
}