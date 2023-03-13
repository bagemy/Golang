package main

import "fmt"

func main() {
	msg := "Hello World-설정완료 한글폰트수정건"
	say(msg)
}

func say(msg string) {
	fmt.Println(msg)
	fmt.Printf("Message변수 주소:%p\n", &msg)
}
