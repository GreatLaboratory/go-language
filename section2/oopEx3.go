package section2

import "fmt"

func OOPEX3() {
	aa := NewAccount1()
	bb := NewAccount2("mg", 150000, 0.0023)
	cc := NewAccount2("ym", 950000, 0.0083)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println()

	CalculateD(*bb)
	CalculateP(cc)

	fmt.Println(bb) // 초기화한 값이랑 차이가 없음
	fmt.Println(cc) // 초기화한 값이랑 차이가 있음. 포인터로 메모리 주소에 접근했기 때문에 다 바뀜
	fmt.Println()

	bb.BonusD(777)
	cc.BonusP(777) // 얘만 값이 바뀜
	fmt.Println()

	fmt.Println(bb)
	fmt.Println(cc)

	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재사용하는 장점
	// 상속이 없는 go언어에서 메소드 상속을 위한 패턴
	emp := Employee{
		name:   "mg",
		salary: 60000,
		bonus:  100,
	}
	exec := Executives{
		Employee: Employee{
			name:   "ym",
			salary: 50000,
			bonus:  100,
		},
		specialBonus: 500,
	}
	fmt.Println(emp.Calculate())
	fmt.Println(exec.Calculate() + exec.specialBonus) // Employee의 리시버함수를 사용할 수 있다.
	fmt.Println(exec.Calculate())                     // 오버라이딩된 메소드의 매개변수 타입을 보고 Executive로 적용하여 메소드가 실행된다.
}

type Account2 struct {
	no       string
	balance  float64
	interest float64
}

// 구조체 생성자 함수
func NewAccount1() *Account2 {
	return &Account2{} // 구조체 인스턴스를 생성한 뒤 포인터를 리턴
}
func NewAccount2(no string, balance, interest float64) *Account2 {
	return &Account2{no, balance, interest}
}

// 매개변수로 포인터 넘기기 비교
func CalculateD(a Account2) {
	a.balance += a.balance * a.interest
}
func CalculateP(a *Account2) {
	a.balance += a.balance * a.interest
}

// 리시버 함수
func (a Account2) BonusD(bonus float64) {
	a.balance += bonus
}
func (a *Account2) BonusP(bonus float64) {
	a.balance += bonus
}

// 상속
type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

// 메소드 오버라이딩
func (e Executives) Calculate() float64 {
	return e.Employee.salary + e.Employee.bonus + e.specialBonus
}

type Executives struct {
	Employee     // is a 관계
	specialBonus float64
}
