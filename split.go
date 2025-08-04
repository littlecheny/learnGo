package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

type node struct{
	value string
	next *node
}

func efil(city string) {
	words := strings.Split(city, " ")
	fmt.Println(words)

	var head *node
	var current *node

	for _, w := range words {
		n := &node{w, nil}
		if head ==nil {
			head = n
			current = n
		}else{
			current.next = n
			current = n
		}
	}

	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var city string
	city, _ = reader.ReadString('\n')
	city = strings.TrimSpace(city)
	efil(city)
}