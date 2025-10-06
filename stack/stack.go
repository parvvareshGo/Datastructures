package main

import "fmt"

// Queue structure
type Queue struct {
	items []int
}

// Enqueue: add item to queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue: remove first item
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

func main() {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Dequeue:", queue.Dequeue()) // 10
	fmt.Println("Dequeue:", queue.Dequeue()) // 20
}
