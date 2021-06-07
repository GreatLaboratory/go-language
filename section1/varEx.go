package section1

import "fmt"

func varEx() {
	// var a int
	// var b string
	// var c, d, e int
	// var f, g, h int = 1, 2, 3
	// var i int = 1
	// var j string = "hi golang"
	// var num = 1
	// var str = "hohoho"
	// var m = true
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)
	name = "updated"
	height = 456
	weight = 33.33
	isRunning = true
	fmt.Println(name, height, weight, isRunning)

	// 짧은 선언
	// 함수 안에서만 사용 가능 (전역x)
	// 선언 후 다시 할당하면 예외 발생
	// 이 함수 안에서 사용되고 끝나면 메모리에서 지워주는 동작
	// if문에서 많이 쓰임
	shortVar1 := 3
	// shortVar1 := 78
	shortVar2 := "hihi"
	shortVar3 := true
	fmt.Println(shortVar1, shortVar2, shortVar3)

	if i := 3; i < 44 {
		fmt.Println("gooed", i)
	}

	// 상수
	// 변수 선언과 동시에 값 할당까지 해야한다. ex) const a string  -> x
	// 함수의 리턴값으로 값 할당이 안된다. ex) const a = getName()  -> x
	const aa = "dfdf"
	const bb, cc int = 3, 4
	const dd, ee, ff = true, 0.25, "df"
	const (
		x, y  = 5, 9
		ii, k = 0.2, 0.639
	)

	// 열거형
	// 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
	// iota는 0부터 1씩 증가된다.
	const (
		_ = iota
		Jan
		Feb
		Mar
		_
		May
		Jun
	)
	fmt.Println(Jan, Feb, Mar, May, Jun) // 1 2 3 5 6
}
