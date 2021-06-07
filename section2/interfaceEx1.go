package section2

import "fmt"

func InterfaceEx1() {
	// 객체의 동작과 골격만 표현
	// 단순히 동작에 대한 방법만 표시
	// 추상화 제공
	// 인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능
	// 덕타이핑 : go언어의 독창적인 특성 - 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드가 있는지로 판단하는 방식
	// 덕타이핑 : 오리처럼 걷고, 소리내고, 헤엄 등의 행동을 하면 오리라고 할 수 있다.
	var t test
	fmt.Println(t) // <nil>

	dog1 := Dog{
		name:   "poll",
		weight: 45,
	}
	dog2 := Dog{
		name:   "marry",
		weight: 15,
	}
	inter1 := AnimalBehavior(dog1)
	inter1.bite()

	inters := []AnimalBehavior{dog1, dog2}
	for _, inter := range inters {
		inter.bite()
	}

	cat1 := Cat{
		name:   "nabi",
		weight: 11,
	}

	// dog1의 타입인 Dog가 전부 구조체 리시버로 act()매개변수로 들어올 인터페이스타입의 메소드를 다 구현해놔야지만 가능
	// Dog가 bite(), sound(), run() 전부 있으면 animalBevior이다. -> 이게 바로 덕타입
	act(dog1)
	act(cat1)

	// 익명 인터페이스
	act2(dog1)
	act2(cat1)

	// any타입을 인터페이스로 만들 수 있다.  -> 거의 만능이라고 함.
	// 함수 매개변수, 리턴값, 구조체 필드 등으로 사용 가능
	x := 1
	y := "str"
	z := true
	printValue(dog1)
	printValue(cat1)
	printValue(cat1)
	printValue(x)
	printValue(y)
	printValue(z)
	printValue([]Dog{})
	printValue([5]Dog{})
	printValue(map[string]int{"aa": 456})
}

// 빈 인터페이스는 nil로 초기화
type test interface {
}

type Dog struct {
	name   string "이름"
	weight int    "몸무게"
}

type Cat struct {
	name   string "이름"
	weight int    "몸무게"
}

func (d Dog) bite() {
	fmt.Println(d.name, ": Dog bites!!")
}
func (d Dog) sound() {
	fmt.Println(d.name, ": Dog barks!!")
}

func (d Dog) run() {
	fmt.Println(d.name, ": Dog is running!!")
}
func (c Cat) bite() {
	fmt.Println(c.name, ": Cat bites!!")
}
func (c Cat) sound() {
	fmt.Println(c.name, ": Cat barks!!")
}

func (c Cat) run() {
	fmt.Println(c.name, ": Cat is running!!")
}

type AnimalBehavior interface {
	bite()
	sound()
	run()
}

func act(animal AnimalBehavior) {
	animal.bite()
	animal.sound()
	animal.run()
}

// 익명 인터페이스
func act2(animal interface{ run() }) {
	animal.run()
}

// any타입을 인터페이스로 만들 수 있다.  -> 거의 만능이라고 함.
func printValue(s interface{}) {
	fmt.Println(s)
}
