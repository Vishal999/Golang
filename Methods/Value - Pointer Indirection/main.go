//Sample program to display methods and pointer indirection

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	//Method with pointer receivers take either a value or a pointer as the receiver
	// when they are called
	//Value semantic used with method with pointer receivers.
	//Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method is
	//pointer receiver.
	v.Scale(10)
	fmt.Println(AbsFunc(v))
	fmt.Println(v.Abs())
	//ScaleFunc(v, 10) //Cpmpiler error!

	p := &Vertex{3, 4}

	//Pointer semantic used with method with pointer receiver
	p.Scale(10)
	ScaleFunc(p, 10)

	fmt.Println(v, p)

	// Method with value receiver take either a value or pointer as receiver when they are
	// called.
	// Pointer semantic used with method with value semantic
	// Go interprets the statement p.Abs() as (*p).Abs()
	fmt.Println(p.Abs())

	//fmt.Println(AbsFunc(p)) //Compiler error
}
