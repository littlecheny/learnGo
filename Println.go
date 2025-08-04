package main

import "fmt"

func life(input string){
	fmt.Println(input)
}

func main (){
	var city string
	fmt.Scanln(&city)
	life(city)
}