package section2

import (
	"fmt"
	"reflect"
	"strconv"
)

func InterfaceEx2() {
	// interface{}로 받으면 any타입이라 이걸 실제 타입으로 변환 후 사용해야할 경우
	// type assertion
	var a interface{} = 15
	b := a
	c := a.(int)
	// d := a.(float64) // 이건 에러가 나옴 - panic: interface conversion: interface {} is int, not float64

	fmt.Println(a, b, c)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

	// 많이 쓰이는 패턴
	// interface{}로 받은 변수읱 타입을 체크
	if v, ok := a.(int); ok {
		fmt.Println(v, ok)
	}
	if v, ok := a.(float64); !ok {
		fmt.Println(v, ok)
	}

	checkType(true)
	checkType(1)
	checkType(1.23)
	checkType("hihi")
	checkType(nil)

	fmt.Println(structToMsg(Account{no: "245-901", balance: 10000000, interest: 0.015}))
	fmt.Println(structToMsg(&Account{no: "245-902", balance: 12000000, interest: 0.035}))
	var user = new(Account)
	user.no = "245-903"
	user.balance = 15000000
	user.interest = 0.025

	fmt.Println(structToMsg(user))
}

// 사실 이게 제일 많이 많이 쓰임
func checkType(arg interface{}) {
	switch arg.(type) {
	case bool:
		fmt.Println("this is bool type ->", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("this is int type ->", arg)
	case float32, float64:
		fmt.Println("this is float type ->", arg)
	case string:
		fmt.Println("this is string type ->", arg)
	case nil:
		fmt.Println("this is nil type ->", arg)
	default:
		fmt.Println("what is this type???", arg)
	}
}

func structToMsg(arg interface{}) string {
	switch arg.(type) {
	case Account:
		o := arg.(Account)
		return "Account : " + o.no + " : " + strconv.FormatFloat((o.balance*o.interest+o.balance), 'f', -1, 64)
	case *Account:
		o := arg.(*Account)
		return "*Account : " + o.no + " : " + strconv.FormatFloat((o.balance*o.interest+o.balance), 'f', -1, 64)
	default:
		return "Error"
	}
}
