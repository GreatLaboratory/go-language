package section4

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func ErrorEx1() {
	// f, err := os.Open("unnamed file") // 일부러 예외 발생
	// if err != nil {
	// 	log.Fatal(err.Error()) // 프로그램을 종료시킨다.
	// }
	// fmt.Println(f.Name)

	//-----------------------------------------------

	if a, err := notZero(1); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(a)
	}
	if a, err := notZero(0); err != nil {
		// log.Fatal(err.Error())
		fmt.Println(err.Error())
	} else {
		fmt.Println(a)
	}

	if result, err := power(0, 4); err != nil {
		fmt.Println("ERROR ::: ", err.Error())
	} else {
		fmt.Println("good process : ", result)
	}
	if result, err := power(4, 4); err != nil {
		fmt.Println("ERROR ::: ", err.Error())
	} else {
		fmt.Println("good process : ", result)
	}
}

func notZero(n int) (str string, err error) {
	if n == 0 {
		str = ""
		err = errors.New("0을 입력했습니다. 에러 발생!!")
	} else {
		str = fmt.Sprintf("Hello Golang : %d", n)
	}
	return str, err
}

func power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0을 넣으면 에러에요")
	}
	return math.Pow(f, i), nil
}
