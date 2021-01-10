package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long.\n", v, len(v))
	default:
		fmt.Printf("%T Type unknown.\n", v)
	}
}

func main() {
	do(25)
	do("Hello")
	do(2.78)
}
