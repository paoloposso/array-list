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

	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

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

func TestRemoveAt(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	err := l.RemoveAt(1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedLength := 2
	if l.Length() != expectedLength {
		t.Errorf("List length is incorrect. Expected: %v, got: %v", expectedLength, l.Length())
	}

	first, _ := l.First()
	expectedFirst := 1
	if *first != expectedFirst {
		t.Errorf("First item is incorrect. Expected: %v, got: %v", expectedFirst, *first)
	}

	last, _ := l.Last()
	expectedLast := 3
	if *last != expectedLast {
		t.Errorf("Last item is incorrect. Expected: %v, got: %v", expectedLast, *last)
	}

	items := l.Filter(func(item int) bool {
		return item == 2
	})

	if len(items) > 0 {
		t.Errorf("Item 2 should have been removed from the list")
	}
}

func TestFilterStruct(t *testing.T) {
	type user struct {
		name string
	}

	l := New[user]()

	l.Append(user{name: "Alice"})
	l.Append(user{name: "Bob"})
	l.Append(user{name: "Charlie"})
	l.Append(user{name: "David"})
	l.Append(user{name: "Eve"})

	existingName := "Alice"

	items := l.Filter(func(item user) bool {
		return item.name == existingName
	})
	if len(items) == 0 || items[0].name != existingName {
		t.Errorf("Expected item %v to exist in the list, but it was not found", existingName)
	}

	nonExistingName := "Zack"
	items = l.Filter(func(item user) bool {
		return item.name == existingName
	})
	if len(items) > 0 {
		t.Errorf("Expected item %v to not exist in the list, but it was found", nonExistingName)
	}
}

func TestClear(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Clear()

	if l.Length() != 0 {
		t.Errorf("List should be empty after calling Clear()")
	}
}

func TestGet(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	item, err := l.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedItem := 2
	if *item != expectedItem {
		t.Errorf("Item is incorrect. Expected: %v, got: %v", expectedItem, *item)
	}
}

func TestGetError(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	_, err := l.Get(3)
	if err == nil {
		t.Errorf("Expected error when trying to get an item at an invalid index")
	}
}

func TestIsEmpty(t *testing.T) {
	l := New[int]()

	if !l.IsEmpty() {
		t.Errorf("List should be empty")
	}

	l.Append(1)

	if l.IsEmpty() {
		t.Errorf("List should not be empty")
	}
}

func TestLength(t *testing.T) {
	l := New[int]()

	if l.Length() != 0 {
		t.Errorf("List length should be 0")
	}

	l.Append(1)

	if l.Length() != 1 {
		t.Errorf("List length should be 1")
	}
}

func TestFirst(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	first, err := l.First()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFirst := 1
	if *first != expectedFirst {
		t.Errorf("First item is incorrect. Expected: %v, got: %v", expectedFirst, *first)
	}
}

func TestFirstError(t *testing.T) {
	l := New[int]()

	_, err := l.First()
	if err == nil {
		t.Errorf("Expected error when trying to get the first item of an empty list")
	}
}

func TestLast(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	last, err := l.Last()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedLast := 3
	if *last != expectedLast {
		t.Errorf("Last item is incorrect. Expected: %v, got: %v", expectedLast, *last)
	}
}

func TestLastError(t *testing.T) {
	l := New[int]()

	_, err := l.Last()
	if err == nil {
		t.Errorf("Expected error when trying to get the last item of an empty list")
	}
}

func TestFilter(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	items := l.Filter(func(item int) bool {
		return item%2 == 0
	})

	expectedLength := 2
	if len(items) != expectedLength {
		t.Errorf("Filtered list length is incorrect. Expected: %v, got: %v", expectedLength, len(items))
	}

	expectedItems := []int{2, 4}
	for i, item := range items {
		if item != expectedItems[i] {
			t.Errorf("Filtered item is incorrect. Expected: %v, got: %v", expectedItems[i], item)
		}
	}
}

func TestFilterEmpty(t *testing.T) {
	l := New[int]()

	items := l.Filter(func(item int) bool {
		return item%2 == 0
	})

	if len(items) != 0 {
		t.Errorf("Filtered list should be empty")
	}
}

func TestFilterStructEmpty(t *testing.T) {
	type user struct {
		name string
	}

	l := New[user]()

	items := l.Filter(func(item user) bool {
		return item.name == "Alice"
	})

	if len(items) != 0 {
		t.Errorf("Filtered list should be empty")
	}
}
