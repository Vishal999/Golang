package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Goodbye")
	}()
	go sayHello()
}

func sayHello() {
	fmt.Println("Hello World")
}
