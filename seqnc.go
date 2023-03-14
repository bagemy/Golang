package main

import "fmt"

func main() {
	// var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3, 4} //배열크기 자동으로
	var multiArray [3][4][5]int   // 정의
	multiArray[0][1][2] = 10      // 사용

	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10

	// println(a1)
	fmt.Println(a3)
	// println(multiArray)

	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]
	fmt.Println(s) //2,3,4 출력
}
