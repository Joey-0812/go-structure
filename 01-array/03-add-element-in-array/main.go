package main

import (
	"fmt"
	"go-structure/01-array/03-add-element-in-array/array"
)

func main() {
	a := array.New(3)

	a.AddLast(2)
	a.AddLast(3)
	fmt.Println(a.Data)
	a.AddFirst(1)

	fmt.Println(a.Data)
}
