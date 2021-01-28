package main

import (
	"fmt"
	go_say_hello "github.com/w4ndry/go-say-hello/v2"
)

func main() {
	name := "Eri"

	fmt.Println(go_say_hello.SayHello(name))
}