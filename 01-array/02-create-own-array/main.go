package main

import (
	"fmt"
	"go-structure/01-array/02-create-own-array/array"
)

func main() {
	a := array.New(5)

	fmt.Println(a)
	fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())
}
