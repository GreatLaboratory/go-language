package section1

import "fmt"

func PointerEx() {
	// 변수는 지역성을 가지고 있는데 포인터는 해당 변수의 주소에 직접 엑세스하기 때문에 어디서든 불러올 수 있다.
	// 주소의 값은 직접 변경 불가능
	// 초기화를 안하면 default값은 nil
	// *는 포인터변수를 선언할 때 타입 앞에 붙히고, &는 포인터변수가 아닌 일반적인 변수의 주소값을 반환시킨다.
	// 포인터변수 앞에 *을 붙히면 해당 주소값에 해당하는 실제값을 반환한다. -> 역참조
	// 포인터변수에는 주소값이 할당 되어야한다.

	// 포인터 선언 / 값을 할당한다. 는 주소값을 할당한다는 뜻이다.
	var a *int            // 선언만
	var b *int = new(int) // 선언 + 초기화
	fmt.Println(a)        // nil
	fmt.Println(b)        // 0xc0000140f0

	i := 7
	fmt.Println(i, &i) // 7 0xc0000140f8

	a = &i
	b = &i
	fmt.Println(a, &i, &a, *a) // 0xc0000140f8 0xc0000140f8 0xc000006030 7
	fmt.Println(b, &i, &b, *b) // 0xc0000140f8 0xc0000140f8 0xc000006038 7

	*a = 999
	fmt.Println(a, &i, &a, *a) // 0xc0000140f8 0xc0000140f8 0xc000006030 999
	fmt.Println(b, &i, &b, *b) // 0xc0000140f8 0xc0000140f8 0xc000006038 999

	var x int = 10
	var y int = 10
	rptc(&x)
	vptc(y)
	fmt.Println("x : ", x, ", y : ", y) // x :  33 , y :  10
}

// 함수, 메소드 호출 시에 매개변수 값을 복사 전달 -> 함수, 메소드 내에서는 원본 값 변경 불가능
// 원본 값 변경 위해서 포인터로 전달
// 특히 크기가 큰 배열일 경우 값 복사 시에 시스템에 부담 -> 이를 포인터 전달로 해결
// 근데 슬라이스나 맵은 그 자체가 참조 전달을 하기에 포인터 필요 없음.
func rptc(a *int) {
	*a = 33
}
func vptc(a int) {
	a = 33
}
