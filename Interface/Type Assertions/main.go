package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	//Since i is of type string, f holds zero value of float64 and no panic is triggered.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//Since i holds a string and we are trying to assign float value, it triggers a panic
	f = i.(float64) //panic
	fmt.Println(f)
}
