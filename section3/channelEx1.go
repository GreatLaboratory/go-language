package section3

import (
	"fmt"
	"time"
)

func ChannelEx1() {
	// 채널은 동기식이다. 채널은 참조 타입이다.
	// 고루틴간의 상호 정보(데이터) 교환 및 싱행 흐름 동기화를 위해 사용
	// 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	// 데이터 전달 자료형 선언 후 사용(지정된 타입만 주고받을 수 있음)
	// interface{} 전달을 통해서 자료형 상관없이 전송 및 수신 가능
	// 값을 전달(복사 후 bool, int등) / 포인터, 슬라이스. 맵 등을 전달할 땐 주의해야함. -> 동기화 사용
	// 멀티 프로세싱 처리에서 교착상태 주의해야함.

	fmt.Println("Main : start --->", time.Now())
	ch := make(chan int) // int형 채널을 선언
	go work1(ch)
	go work2(ch)
	num1 := <-ch // 1이나 2가 올 때까지 기다렸다가 수신이 완료되면 다음 코드로 넘어감. ->동기화!!
	num2 := <-ch
	fmt.Println(num1, num2)
	fmt.Println("Main : end --->", time.Now())
}

func work1(ch chan int) { // int형만 채널을 왔다갔다할 수 있다.
	fmt.Println("work1 : start ----> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work1 : end ----> ", time.Now())
	ch <- 1 // 데이터 1을 채널을 통해 보낸다. -송신
}
func work2(ch chan int) {
	fmt.Println("work2 : start ----> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 : end ----> ", time.Now())
	ch <- 2
}
