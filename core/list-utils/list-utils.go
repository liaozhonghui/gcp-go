package main

import (
	"container/list"
	"fmt"
)

type rectangle struct {
	Width, Heigh int
}

func main() {
	var l list.List
	elem := rectangle{Width: 10, Heigh: 20}
	l.PushBack(elem)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("e.value: %+v\n", e.Value)
	}
}
