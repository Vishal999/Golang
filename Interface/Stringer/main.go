package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {

	//Sprintf formats according to a format specifier and returns the resulting string.
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	a := Person{"Vishal", 24}
	b := Person{"Gaurav", 22}
	fmt.Println(a, b)
}
