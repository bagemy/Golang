package main
import "fmt"

func main() {
	var a int
	var b int

	a = 1
	b = 1

	IncVal(a)
	IncPointer(&b)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

func IncVal(x int) {
	x++
	fmt.Println("x : ", x)
}

func IncPointer(x *int) {
	*x = *x + 1
}
