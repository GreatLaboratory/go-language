package section1

import "fmt"

func FunctionEx2() {
	fmt.Println(multiply_sum(1, 2, 3, 4, 5)) // 120 15
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println(multiply_sum(slice...)) // 30240 40 -> n ...int 을 슬라이스에서 뽑아낼 때엔 뒤에 ...을 붙히면 된다.

	// 변수에 함수 할당
	a := []func(int, int) int{plus, minus}
	fmt.Println(a[0](7, 3))
	fmt.Println(a[1](7, 3))

	var f1 func(int, int) int = plus
	f2 := minus
	fmt.Println(f1(7, 3))
	fmt.Println(f2(7, 3))

	m := map[string]func(int, int) int{
		"plus":  plus,
		"minus": minus,
	}
	fmt.Println(m["plus"](7, 3))
	fmt.Println(m["minus"](7, 3))

	fmt.Println(factorial(5)) // 120

	// 익명함수
	func() {
		fmt.Println("hi")
	}()
	func(m string) {
		fmt.Printf("hi %s\n", m)
	}("great")
	func(x, y int) {
		fmt.Println("sum : ", x+y)
	}(10, 20)
	result := func(x, y int) int {
		return x * y
	}(10, 20)
	fmt.Println("tot : ", result)
}

// 가변 인자 -> 매개 변수의 개수가 동적으로 변하는게 가능하다.
func multiply_sum(n ...int) (tot, sum int) {
	tot = 1
	sum = 0
	for _, item := range n {
		tot *= item
		sum += item
	}
	return tot, sum
}

func plus(x, y int) int {
	return x + y
}
func minus(x, y int) int {
	return x - y
}

// 재귀함수
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
