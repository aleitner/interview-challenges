package main

import "fmt"

type Thing struct {
	data string
	next *Thing
}

type Queue struct {
	current *Thing
	last *Thing
}

func (q *Queue) Insert(data string) {
	if q.current == nil {
		q.current = &Thing{
			data: data,
		}

		return
	}

	var current *Thing
	var next *Thing
	current = q.current
	next = current.next
	for {
		if next == nil {
			current.next = &Thing{
				data: data,
			}
			return
		}

		current = next
		next = next.next
	}
}

func (q *Queue) Pop() (string, error) {
	poppedItem := q.current

	if poppedItem == nil {
		return "", fmt.Errorf("No data on queue")
	}

	q.current = q.current.next
	return poppedItem.data, nil
}

func main() {

}