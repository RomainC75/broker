package service

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	Start   *QueueElement[T]
	End     *QueueElement[T]
	Size    int
	MaxSize int
}

type QueueElement[T any] struct {
	Previous *QueueElement[T]
	Next     *QueueElement[T]
	Content  T
}

func NewQueue[T any](maxSize int) *Queue[T] {
	return &Queue[T]{
		Start:   nil,
		End:     nil,
		Size:    0,
		MaxSize: maxSize,
	}
}

func (q *Queue[T]) AddElement(content T) (*QueueElement[T], error) {
	newElement := QueueElement[T]{
		Content: content,
	}
	if q.Size == q.MaxSize {
		return nil, errors.New("queue is full")
	}
	if q.Size == 0 {
		q.Size = 1
		q.Start = &newElement
		q.End = &newElement
		return &newElement, nil
	}
	q.End.Next = &newElement
	previous := q.End
	q.End = &newElement
	newElement.Previous = previous
	q.End = &newElement
	q.Size += 1
	return &newElement, nil
}

func (q *Queue[T]) DeleteElement(index int) (*QueueElement[T], error) {
	fmt.Println("deletej")
	if index < 0 || index >= q.MaxSize {
		return nil, errors.New("cannot delete this element")
	}
	el := q.Start
	for index > 0 {
		el = el.Next
		index--
	}
	el.Next.Previous, el.Previous.Next = el.Previous.Next, el.Next
	return el, nil
}

func (q *Queue[T]) GetValues() []T {
	list := make([]T, 0, q.MaxSize)
	el := q.Start
	for el.Next != nil {
		list = append(list, el.Content)
		el = el.Next
	}
	list = append(list, el.Content)
	return list
}

func (q *Queue[T]) DisplayElements() {
	el := q.Start
	// for i := 0; i < q.Size; i++ {
	for el.Next != nil {
		fmt.Println(el.Content)
		el = el.Next
	}
	fmt.Println(el.Content)
}

func (q *Queue[T]) ExtractFirstValue() (T, error) {
	if q.Size > 0 {
		value := q.Start
		q.Size--
		if q.Size == 0 {
			q.End = nil
			q.Start = nil
		} else {
			q.Start = value.Next
		}
		return value.Content, nil
	} else {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
}
