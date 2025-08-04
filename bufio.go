package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func efil(city string) {
	fmt.Println(city)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var city string
	city, _ = reader.ReadString('\n')
	city = strings.TrimSpace(city)
	efil(city)
}