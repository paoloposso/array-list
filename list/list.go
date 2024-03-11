package list

import (
	"errors"
	"fmt"
	"strings"
)

type list[T any] struct {
	array *[]T
}

type List[T any] interface {
	Length() int
	IsEmpty() bool
	First() (*T, error)
	Last() (*T, error)
	Tail() List[T]
	String() string
	Get(index int) (*T, error)
	Append(item T)
	Prepend(item T)
}

// New returns an empty list
func New[T any]() List[T] {
	return &list[T]{
		array: &[]T{},
	}
}

// Size returns the size of the list
func (l *list[T]) Length() int {
	return len(*l.array)
}

// IsEmpty returns true if the list contains zero elements and false otherwise
func (l *list[T]) IsEmpty() bool {
	return l.array == nil || len(*l.array) == 0
}

// First returns the first element of the list or an error if the list is empty
func (l *list[T]) First() (*T, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	return &(*l.array)[0], nil
}

// Last returns the last element of the list or an error if the list is empty
func (l *list[T]) Last() (*T, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	return &(*l.array)[len(*l.array)-1], nil
}

// Tail returns all elements of the list except for the first (head)
func (l *list[T]) Tail() List[T] {
	if l.IsEmpty() || len(*l.array) == 1 {
		return &list[T]{array: &[]T{}}
	}
	a := (*l.array)[1:]
	return &list[T]{
		array: &a,
	}
}

// String returns a string representation of the list.
func (l *list[T]) String() string {
	var allElements []string
	for i := 0; i < len(*l.array); i++ {
		allElements = append(allElements, fmt.Sprintf("%v", (*l.array)[i]))
	}
	return fmt.Sprintf("[%s]", strings.Join(allElements, " "))
}

// Get returns the item at the provided index, or an error if the index is out of bounds
func (l *list[T]) Get(index int) (*T, error) {
	if index > len(*l.array)-1 {
		return nil, errors.New("index out of bounds")
	}
	return &(*l.array)[index], nil
}

// Append adds a new element to the end of the list
func (l *list[T]) Append(item T) {
	*l.array = append(*l.array, item)
}

// Prepend adds a new element to the beginning of the list
func (l *list[T]) Prepend(item T) {
	newArray := make([]T, len(*l.array)+1)
	newArray[0] = item
	copy(newArray[1:], *l.array)
	l.array = &newArray
}

// RemoveAt removes the element at the provided index from the list
func (l *list[T]) RemoveAt(index int) error {
	if index < 0 || index >= len(*l.array) {
		return errors.New("index out of bounds")
	}
	*l.array = append((*l.array)[:index], (*l.array)[index+1:]...)
	return nil
}
