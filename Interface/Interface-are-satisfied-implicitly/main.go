//Sample Program to depict implicit interface implementation

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means that type T implements interface I.
// but we don't need to explicitly declare that.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello, World"}
	i.M()
}
