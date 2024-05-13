package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

// method Enqueue add new items to Queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// method Dequeue deletes the first item in the Queue
func (q *Queue) Dequeue() int {

	firstItem := q.items[0]
	q.items = q.items[1:]
	return firstItem
}

func main() {
	fmt.Printf("This is a Program for showcasing Queue Data stuctures\n\n")

	myQueue := Queue{}

	fmt.Println("Adding values to the Queue -> ")

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println("Queue Contains : ", myQueue)

	fmt.Println("Delete a value from Queue ->")

	myQueue.Dequeue()

	fmt.Println("Queue after Dequeue operation: ", myQueue)
}
