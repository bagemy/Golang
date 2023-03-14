package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}
type dict struct {
	data map[int]string
}

//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

//Rect - struct 정의
type Rect struct {
	width, height int
}

//Rect의 area() 메소드
func (r Rect) area() int {
	return r.width * r.height
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}

	fmt.Println(p1, p2)

	dic := newDict() // 생성자 호출
	dic.data[1] = "A"

	rect := Rect{10, 20}
	area := rect.area() //메서드 호출
	println(area)

	rect1 := Rect{100, 200}
	area1 := rect1.area2() //메서드 호출
	println(area1)
}
