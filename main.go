package structures

import "github.com/paoloposso/go-structures/list"

// NewList creates a new list
func NewList[T any]() list.List[T] {
	return list.New[T]()
}
