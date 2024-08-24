package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Values = []int{1, 2, 3, 4, 5, 6, 7, 8}

const IndexToRemove = 3

var ResultValues = []int{1, 2, 3, 5, 6, 7, 8}

var FirstOut = 1
var FirstOutArray = []int{2, 3, 5, 6, 7, 8}

func TestLinkedList(t *testing.T) {
	values := Values
	linkedList := NewQueue[int](10)
	for _, value := range values {
		linkedList.AddElement(value)
	}

	createdValues := linkedList.GetValues()
	fmt.Println(createdValues)
	assert.Equal(t, values, createdValues, "the slicees should be equals")

	// * Add Value
	newValue := 9
	values = append(values, newValue)
	_, err := linkedList.AddElement(newValue)
	assert.NoError(t, err, "not supposed to raise an error")
	valuesWithAdded := linkedList.GetValues()
	assert.Equal(t, values, valuesWithAdded, "new slice not correct")

	// * Remove value at index
	index := 3
	deleteElementListRef := append(values[:index], values[index+1:]...)
	_, err = linkedList.DeleteElement(index)
	assert.NoError(t, err, "not supposed to raise an error")
	deleteElementList := linkedList.GetValues()
	assert.Equal(t, deleteElementListRef, deleteElementList, "should have the same element guy")

	// * First out
	firstOut, err := linkedList.ExtractFirstValue()
	fmt.Println("-> firstOut : ", firstOut)
	assert.NoError(t, err, "not supposed to raise an error with this ExtractFirstValue()")
	assert.Equal(t, firstOut, FirstOut)

	// * First out when Queue is only one element
	addedElement := 2
	emptyQueue := NewQueue[int](10)
	emptyQueue.AddElement(addedElement)
	extractedValue, err := emptyQueue.ExtractFirstValue()
	assert.NoError(t, err, "should be an error when extracting an empty Queue")
	assert.Equal(t, extractedValue, addedElement)

	// * First out when empty Queue
	_, err = emptyQueue.ExtractFirstValue()
	assert.Error(t, err, "should be an error when extracting an empty Queue")
	assert.EqualError(t, err, "queue is empty")

}
