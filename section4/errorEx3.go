package section4

import (
	"fmt"
	"os"
)

func ErrorEx3() {
	// Panic Recover
	// 자바의 try catch랑 비슷
	// 사용자가 에러 발생 가능

	// Panic함수는 호출 즉시, 해당 메소드를 즉시 중단시키고 defer함수를 호출하고 자기자신을 호출한 곳으로 리턴
	// 런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생시킬 때 중요하다.
	// 문법적인 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

	// Recover함수는 Panic으로 발생한 에러를 복구 후 코드를 재실행(프로그램 종료되지 않는다.)
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// Panic에서 설정한 메세지를 받아올 수 있다.

	fmt.Println("Start Main")
	// panic("Error occurred : user stopped!") // 방법1
	// log.Panic("Error occurred : user stopped!") // 방법2

	// runFunc()
	openFile("goog.txt")
	fmt.Println("End Main")
}

func runFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error Message : ", r)
		}
	}()
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 10; i++ {
		fmt.Println(a[i]) // (자연스러운 패닉)
	}
	panic("Error occurred!") // 패닉이 발생하면 즉시 해당 함수의 defer함수가 있는지를 찾고 그걸 실행시킨다. (개발자가 만든 고의적인 패닉)
	fmt.Println("ㅇㅇㅇㅇ")      // 따라서 해당 라인은 실행이 안된다.
}

func openFile(fileName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()
	if f, err := os.Open(fileName); err != nil {
		panic(err)
	} else {
		fmt.Println("file name : ", f.Name())
		defer f.Close()
	}
}
