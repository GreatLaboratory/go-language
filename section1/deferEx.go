package section1

import "fmt"

func DeferEx() {
	// defer 지연함수 - 미뤄놨다가 나중에 하는 함수
	// defer를 호출한 함수가 종료되기 직전에 호출된다.
	// finally랑 비슷
	// db작업 리소스 끝나고 종료, 열린 파일 닫기, mutex 잠금 해제
	f1()

	stack() // 스택의 구조로 10부터 출력된다.

	a()
	// start : b
	// hi
	// end : b
}

// defer
func f1() {
	fmt.Println("start!!")
	defer f2() // f1함수가 끝나기 직전인 마지막에 호출된다.
	fmt.Println("end!!")
}
func f2() {
	fmt.Println("called!!")
}

// defer의 스택구조
func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}

// defer 이후 함수 중첩문
// defer는 무조건 바로 뒤에 함수인 end함수만 마지막으로 한다. 따라서 start함수가 가장 먼저 실행된다.
// 왠만하면 defer 뒤에는 중첨함수를 쓰지않는 것을 권장
func a() {
	defer end(start("b"))
	fmt.Println("hi")
}
func start(m string) string {
	fmt.Println("start : ", m)
	return m
}
func end(m string) {
	fmt.Println("end : ", m)
}
