package main

import "fmt"

func main() {
	var a int
	//변수, 이름, type(int)
	var msg string
	//변수, 이름, type(string)
	a = 20               // a 라는 변수에 20
	msg = "Good Morning" //msg 라는 변수에 "Good Morning"

	fmt.Println("--------------------------")

	var b int = 35
	var msg123 string = "반갑습니다"
	// 변수의 값 변경
	b = 55
	msg123 = "수고하셨어요"
	fmt.Println(msg, a)
	fmt.Println(msg123, b)

	fmt.Println("--------------------------")
	c := 23        // var a int = 23 -> a:=23
	d := "나는 누굴까요" // var d string = "" -> d:= ""
	fmt.Println(c, d)

	fmt.Println("--------------------------")
	e := 3.14
	var f = "나는 3.14이다"

	fmt.Println(f, e)
}
