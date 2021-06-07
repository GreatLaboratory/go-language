package section1

import "fmt"

func ClosureEx() {
	// 클로저는 함수 밖에 있는 변수를 캐치해서 접근 사용
	// 익명함수를 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에 함수가 실행될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 목적 : 함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서
	// 무분별한 전역변수 남발 방지

	multiply := func(x, y int) int { // 익명함수를 변수에 할당
		return x * y
	}
	r1 := multiply(5, 10)
	fmt.Println(r1) // 50

	m, n := 5, 10 // 이 변수들이 스냅샷으로 캡쳐되었다.
	sum := func(x int) int {
		// m, n := 5, 10 -> 이런 식으로 m과 n이 내부에서 선언되어야하는데 외부에서 선언되고 있다.
		return m + n + x // 지역변수가 소멸되지 않는다. (함수 호출 시마다 사용가능)
	}
	r2 := sum(5)
	fmt.Println(r2) // 20

	cnt := increaseCnt()
	fmt.Println(cnt()) // 1
	fmt.Println(cnt()) // 2
	fmt.Println(cnt()) // 3
	fmt.Println(cnt()) // 4
	fmt.Println(cnt()) // 5
}
func increaseCnt() func() int {
	n := 0 // increaseCnt 함수가 끝나면 garbage collection에 들어가야하는데 스냅샷으로 찍혀서 계속 값이 바뀐다.
	return func() int {
		n++
		return n
	}
}
