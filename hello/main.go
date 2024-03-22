package main

import (
	"fmt"

	"rsc.io/quote"
	// values "github.com/KaiAlan/go-basics"
)

func forlooops() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func closures() {
	const name string = "Alen"
	var i int = 9

	if i > 10 {
		fmt.Println(name)
	} else if i == 10 {
		fmt.Println(name + " is a closure function")
	} else {
		fmt.Println("This is a closure function with if-alse statements")
	}
}

func main() {
	fmt.Println("Hello from kai")
	fmt.Println(quote.Glass())

	// values.Values()
	Variables()

	x := forlooops

	x()

	closures()

	// SwitchCases()

	// Maps()

	// Slices()
}
