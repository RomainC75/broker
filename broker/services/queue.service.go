package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// * 2 states :
// * wainting
// * handking
// * done = deleted

type Queue[T any] struct {
	Start         *QueueElement[T]
	End           *QueueElement[T]
	size          int
	MaxSize       int
	HandlingTasks []ProcessingTask[T]
}

type QueueElement[T any] struct {
	Previous *QueueElement[T]
	Id       uuid.UUID
	Next     *QueueElement[T]
	Content  T
}

type ProcessingTask[T any] struct {
	Id      uuid.UUID
	Content T
	Date    time.Time
}

func NewQueue[T any](maxSize int) *Queue[T] {
	return &Queue[T]{
		Start:   nil,
		End:     nil,
		size:    0,
		MaxSize: maxSize,
	}
}

func (q *Queue[T]) AddElement(content T) (*QueueElement[T], error) {
	newElement := QueueElement[T]{
		Content: content,
	}
	if q.size == q.MaxSize {
		return nil, errors.New("queue is full")
	}
	if q.size == 0 {
		q.size = 1
		q.Start = &newElement
		q.End = &newElement
		return &newElement, nil
	}
	q.End.Next = &newElement
	previous := q.End
	q.End = &newElement
	newElement.Previous = previous
	q.End = &newElement
	q.size += 1
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

func (q *Queue[T]) GetFirstValueAndToHandling() (T, uuid.UUID, error) {
	if q.size > 0 {
		value := q.Start
		q.size--
		if q.size == 0 {
			q.End = nil
			q.Start = nil
		} else {
			q.Start = value.Next
		}
		id := q.AddHandlingTask(value.Content)
		return value.Content, id, nil
	} else {
		var zeroValue T
		return zeroValue, uuid.New(), errors.New("queue is empty")
	}
}

func (q *Queue[T]) ReadFirstContent() (T, error) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("empty queue")
	}
	return q.Start.Content, nil
}

func (q *Queue[T]) GetSize() int {
	return q.size
}

// =====Handling Queue====== //

// TODO : use map instead of slice !!

func newProcesssingElement[T any](content T) ProcessingTask[T] {
	return ProcessingTask[T]{
		Id:      uuid.New(),
		Date:    time.Now(),
		Content: content,
	}
}

func (q *Queue[T]) AddHandlingTask(content T) uuid.UUID {
	pe := newProcesssingElement(content)
	q.HandlingTasks = append(q.HandlingTasks, pe)
	return pe.Id
}

func (q *Queue[T]) RemoveHandlingTask(id uuid.UUID) (T, error) {
	for index, ht := range q.HandlingTasks {
		if ht.Id.String() == id.String() {
			q.HandlingTasks = append(q.HandlingTasks[:index], q.HandlingTasks[index+1:]...)
			return ht.Content, nil
		}
	}
	var zeroTValue T
	return zeroTValue, errors.New("not found")
}
