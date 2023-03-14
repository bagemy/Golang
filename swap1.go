package main

import "fmt"

func main() { 
	a := 1
	b := 2
	fmt.Printf("a = %d, b = %d\n", a, b)
	swap(a, b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(a int, b int) {
	temp := a
	a = b
	b = temp
}
