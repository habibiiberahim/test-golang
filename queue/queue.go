package main

import "fmt"

type UniqueQueue struct {
	data   []interface{}
	length int
}

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) *UniqueQueue {
	return &UniqueQueue{
		data: make([]interface{}, size),
	}
}

func (q *UniqueQueue) Push(key interface{}) {
	// checking queue length equal with queue size
	if q.length == len(q.data) {
		q.Pop()
	}
	q.data[q.length] = key
	q.length++
}
func (q *UniqueQueue) Pop() interface{} {
	//if queue length small equals zero queue is empty
	if q.length <= 0 {
		fmt.Println("can't pop empty queue")
	}
	popData := q.data[0]
	//move data to the front queue
	for i := 0; i < q.length-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.length--
	return popData

}

func (q *UniqueQueue) Len() int {
	return q.length
}

func (q *UniqueQueue) Contains(key interface{}) bool {
	contain := false
	//check data in queue with key
	for i := 0; i < q.length; i++ {
		if q.data[i] == key {
			contain = true
		}
	}
	return contain
}

func (q *UniqueQueue) Keys() []interface{} {
	//return data queue
	return q.data[0:q.length]
}
