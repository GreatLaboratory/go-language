package section1

import (
	"fmt"
	"strconv"
)

func FunctionEx1() {
	// 매개변수로 함수 전달
	fmt.Println(sum(4, 8))
	fmt.Println(strconv.Itoa(sum(4, 8))) // toString
	sum2(100, add)                       // 115

	// 그냥 값과 참조값의 차이
	num := 7
	multi_value(num)
	fmt.Println(num) // 7
	multi_reference(&num)
	fmt.Println(num) // 70

	// 반환 값이 다수 가능
	// x, _ := plus_minus(7, 3) -> 이렇게도 가능하다.
	x, y := plus_minus(7, 3)
	fmt.Println(x, y) // 10 4
}

func sum(x int, y int) int {
	return x + y
}
func sum2(i int, f func(int, int)) {
	f(i, 15)
}
func add(x, y int) {
	fmt.Println(x + y)
}

func multi_value(i int) {
	i *= 10
}
func multi_reference(i *int) {
	*i *= 10
}

func plus_minus(x, y int) (int, int) {
	return x + y, x - y
}
func plus_minus2(x, y int) (plus, minus int) {
	plus = x + y
	minus = x - y
	return plus, minus
}
