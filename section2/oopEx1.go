package section2

import "fmt"

func OOPEX1() {
	// 클래스와 상속 개념이 없음.
	// 인터페이스 -> 다형성 지원 가능
	// 구조체를 통한 클래스 형태 개발 가능
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본 타입, 함수 등을 전부 사용자가 정의할 수 있다.
	// 구조체와 메소드는 바인딩된다.
	// 맵과 슬라이스는 기본적으로 참조 전달이다. 구조체도 참조 전달이다.

	car1 := Car{
		name:  "bmw",
		color: "black",
		price: 500000,
		tax:   100000,
	}
	car2 := Car{
		name:  "benz",
		color: "white",
		price: 500000,
		tax:   100000,
	}
	fmt.Println(car1, &car1) // 이렇게 &를 사용하면 원래 주소가 나오는데 사용자 정의 구조체에선 안 나온다.
	fmt.Println(car2, &car2)
	fmt.Printf("%p   %p\n", &car1, &car2) // 이렇게 해야 나옴.

	fmt.Println(price(car1), price(car2))
	fmt.Println(car1.price2(), car1.price2())

	a := cnt(15)
	var b cnt = 15
	conversionA(int(a)) // cnt가 아닌 무조건 int형이어야해서 형변환 없이 그냥 a를 넣으면 에러다.
	conversionB(b)

	var orderPrice totCost
	orderPrice = func(cnt, price int) int {
		return cnt*price + 10000
	}
	describe(80, 100, orderPrice)

	bs1 := shoppingBasket{
		cnt:   4,
		price: 500,
	}
	fmt.Println(bs1.purchase())
	bs1.addPurchase1(1, 700) // 여기선 참조전달이므로 주소값에 접근하여 bs1의 속성들의 값을 바꾼다. (원본 값 수정o)
	bs1.addPurchase2(1, 700) // 여기선 참조전달이 아니라 bs1의 속성들의 값을 바꾸지 못한다. (원본 값 수정x)
	fmt.Println(bs1.purchase())
}

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

// 이건 그냥 일반 함수
func price(c Car) int64 {
	return c.price + c.tax
}

// 이게 구조체와 바인딩된 메소드
func (c Car) price2() int64 {
	return c.price + c.tax
}

// 사용자 정의 기본 타입
type cnt int

func conversionA(i int) {
	fmt.Println("Default Type : ", i)
}
func conversionB(i cnt) {
	fmt.Println("Custom Type : ", i)
}

// 사용자 정의 함수 타입
type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, orderPrice : %d\n", cnt, price, fn(cnt, price))
}

// 매개변수 - 원본수정(참조 형식) vs 원본수정x(값 전달 형식)
type shoppingBasket struct {
	cnt   int
	price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}
func (b *shoppingBasket) addPurchase1(cnt, price int) {
	b.cnt += cnt
	b.price += price
	// *b.cnt += cnt 이런 식으로 역참조하는걸 스킵할 수 있다.
}
func (b shoppingBasket) addPurchase2(cnt, price int) {
	b.cnt += cnt
	b.price += price
}
