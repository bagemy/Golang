package  main

import "fmt"

func swap(x *int, y *int) {
    temp  := *x
    *x = *y
    *y = temp
}

func main() {
    a := 10
    b := 20
    fmt.Println("before",a,b)
    swap (&a, &b)
    fmt.Println("after" ,a,b)
}
