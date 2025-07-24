package main

import "fmt"

func main() {
	// Arrays
	var arr [5]int
	arr[0] = 100
	arr[4] = 90
	fmt.Println(arr)

	// Slices
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// Maps
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	// Structs
	type person struct {
		name string
		age  int
	}
	p := person{name: "John", age: 30}
	fmt.Println(p)
}
