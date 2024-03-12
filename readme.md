# Go Generic List Package

This package provides a generic List data structure for Go. It uses Go 1.18's generics feature, so it requires Go 1.18 or later.

## Features

- Create a new list with `New()`
- Get the length of the list with `Length()`
- Check if the list is empty with `IsEmpty()`
- Get the first item in the list with `First()`
- Get the last item in the list with `Last()`
- Get all items in the list except the first one with `Tail()`
- Get a string representation of the list with `String()`
- Get the item at a specific index with `Get(index int)`
- Add an item to the end of the list with `Append(item T)`
- Add an item to the beginning of the list with `Prepend(item T)`
- Remove the item at a specific index with `RemoveAt(index int)`
- Get a new list containing only the items that satisfy a specific condition with `Filter(fn func(T) bool)`
- Remove all items from the list with `Clear()`

## Usage

```go
import "list"

func main() {
    l := list.New[int]()
    l.Append(1)
    l.Append(2)
    l.Append(3)
    fmt.Println(l) // Outputs: [1 2 3]
}
```

In this example, we create a new list of integers, append some numbers to it, and then print it.

## Error Handling

Some methods, like `First()`, `Last()`, and `Get(index int)`, return an error if the operation cannot be performed (for example, if the list is empty or the index is out of bounds). Be sure to check for these errors in your code.

## Note

This package uses Go's new generics feature, which is available in Go 1.18 and later. If you're using an earlier version of Go, you'll need to upgrade to use this package.