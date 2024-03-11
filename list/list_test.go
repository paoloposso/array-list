package list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	l := New[int]()

	// Append some items
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Check if the last item is correct
	last, err := l.Last()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedLast := 3
	if *last != expectedLast {
		t.Errorf("Last item is incorrect. Expected: %v, got: %v", expectedLast, *last)
	}
}

func TestPrepend(t *testing.T) {
	l := New[int]()

	// Prepend some items
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	// Check if the first item is correct
	first, err := l.First()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFirst := 3
	if *first != expectedFirst {
		t.Errorf("First item is incorrect. Expected: %v, got: %v", expectedFirst, *first)
	}
}

func TestUpdateAsReference(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	appendOneItem(l, 4)

	// Check if the last item is updated
	last, _ := l.Last()
	expectedLast := 4
	if *last != expectedLast {
		t.Errorf("First item is incorrect. Expected: %v, got: %v", expectedLast, *last)
	}
}

func appendOneItem(l List[int], value int) {
	l.Append(value)
}

func TestTailRecursion(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	result := tailRecursion(l)

	expectedResult := 10
	if result != expectedResult {
		t.Errorf("Tail recursion result is incorrect. Expected: %v, got: %v", expectedResult, result)
	}
}

func tailRecursion(l List[int]) int {
	// Base case: empty list
	if l.IsEmpty() {
		return 0
	}

	first, _ := l.First()

	return *first + tailRecursion(l.Tail())
}
