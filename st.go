package main

import (
	"fmt"
	"strings"
)

func main() {
	command := "http://www.example.com/page1.html 192.168.2.3/user.host.name jabroni GET"
	command1 := strings.Split(command, " ")
	fmt.Println("0-- ", command1[0])
	fmt.Println("1-- ", command1[1])
	fmt.Println("2-- ", command1[2])
	fmt.Println("3-- ", command1[3])

}
