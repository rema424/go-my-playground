package main

import (
	"app/lib"
	"fmt"
)

func main() {
	d1 := map[string]interface{}{
		"name": "Michael",
		"age":  25,
	}

	fmt.Println("Map")
	lib.TypeAsertion(d1)

	fmt.Println("Map Pointer")
	lib.TypeAsertion(&d1)

	d2 := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println("Array")
	lib.TypeAsertion(d2)

	fmt.Println("Array Pointer")
	lib.TypeAsertion(&d2)

	d3 := []int{10, 100, 1000, 10000}

	fmt.Println("Slice")
	lib.TypeAsertion(d3)

	fmt.Println("Slice Pointer")
	lib.TypeAsertion(&d3)

	type User struct {
		Name string
		Age  int
	}
	d4 := User{Name: "Michael", Age: 25}

	fmt.Println("Struct")
	lib.TypeAsertion(d4)

	fmt.Println("Struct Pointer")
	lib.TypeAsertion(&d4)
}
