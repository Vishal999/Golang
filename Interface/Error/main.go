package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

//implementing error interface
func (err MyError) Error() string {
	return fmt.Sprintf("at %v, %q", err.When, err.What)
}

func run() error {
	return &MyError{
		When: time.Now(),
		What: "mission abort, tank has exploded!",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
