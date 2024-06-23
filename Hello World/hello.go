// 패키지 선언
package main

// 프로그램에서 패키지를 가져오는데 사용되며 fmt 패키지는
// 함수로 형식화된 입력/출력을 구현하는 데 사용된다.
import "fmt"

// func: Go언어에서 함수를 만드는데 사용된다.
// main: Go언어의 주요 함수로 매개변수가 없고 아무것도 반환하지 않으며
// 프로그램 실행 시 호출된다.
func main() {
	//Println()이 메소드는 fmt 패키지에 있으며 Println은 인쇄 라인을 의미한다.
	fmt.Println("Hello World")
}
