package main

import "fmt"

func main() {
	// Using var
	var age int32 = 37
	var isCool = true
	isCool = false


	name := "Brad"

	fmt.Println(name, age)
	fmt.Printf("%T\n", age)
	fmt.Println(isCool)
}
