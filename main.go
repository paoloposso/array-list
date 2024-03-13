package structures

import "github.com/paoloposso/go-structures/list"

func NewList[T any]() list.List[T] {
	return list.New[T]()
}
