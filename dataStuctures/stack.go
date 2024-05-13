package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

// method Push add new items to stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// method Pop deletes the last item in the stack
func (s *Stack) Pop() int {
	length := len(s.items)

	lastItem := s.items[length-1]
	s.items = s.items[:length-1]
	return lastItem
}

func main() {
	fmt.Printf("This is a Program for showcasing Stack Data stuctures\n\n")

	myStack := Stack{}

	fmt.Println("Adding values to the Stack -> ")

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println("Stack Contains : ", myStack)

	fmt.Println("Popping a valur from stack ->")

	myStack.Pop()

	fmt.Println("Stack after Pop operation: ", myStack)
}
