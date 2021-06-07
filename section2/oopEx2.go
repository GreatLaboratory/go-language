package section2

import (
	"fmt"
	"reflect"
)

func OOPEX2() {
	a1 := Account{
		no:       "mg",
		balance:  5000,
		interest: 0.037,
	}
	a2 := Account{
		no:      "mg",
		balance: 5000,
	}
	fmt.Println(a1, a1.Calculate())
	fmt.Println(a2) // interest값은 할당이 안되면 float64의 초기값으로 초기화된다. 0

	// 다양한 선언방법
	// &struct, struct : &struct 포인터를 받아오고 역참조를 또 하기 때문에 속도는 조금 느리ㅏㄷ.
	// 인터페이스 메소드를 선언만 해둔 후, 오버라이딩해서 메소드에 포인트 리시버를 사용할 경우 반드시 &struct를 사용해야한다.
	var kim *Account = new(Account)
	kim.no = "kim"
	kim.balance = 12000
	kim.interest = 0.021

	hong := &Account{
		no:       "hong",
		balance:  5000,
		interest: 0.037,
	}

	lee := new(Account)
	lee.no = "good"
	lee.balance = 12000
	lee.interest = 0.021
	fmt.Println(kim)
	fmt.Println(hong)
	fmt.Println(lee)

	// 익명 구조체
	car1 := struct{ name, color string }{"bmw", "red"}
	fmt.Println(car1)

	cars := []struct{ name, color string }{{"bmw", "red"}, {"benz", "black"}, {"kia", "white"}}
	for _, car := range cars {
		fmt.Println(car)
	}

	// 구조체 필드에 태그
	tag := reflect.TypeOf(Person{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

	// 중첩 구조체
	person1 := Person{
		name:    "mg",
		age:     26,
		address: "hanam",
		detail: spec{
			height: 181,
			weight: 78,
		},
	}
	fmt.Println(person1)
	fmt.Println(person1.detail.height, person1.detail.weight)
}

type Account struct {
	no       string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + a.balance*a.interest
}

// 구조체 필드에 태그를 붙힐 수 있다. 주석처럼
type Person struct {
	name    string "이름"
	age     int32  "나이"
	address string "주소"
	detail  spec   "상세"
}

// 중첩 구조체
type spec struct {
	weight int16 "몸무게"
	height int16 "키"
}
