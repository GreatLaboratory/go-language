package section3

import (
	"fmt"
	"time"
)

func ChannelEx4() {
	// 함수 등의 매개변수에 수신 및 발신 전용 채널 지정 가능
	// 전용 채널 설정 후 방향이 다를 경우 예외 발생
	// 발신 전용 channel <- 데이터형
	// 수신 전용 <- channel
	// 매개변수를 통해서 전용 채널 확인할 수 있다.
	// 채널 또한 함수의 반환 값으로 사용 가능

	ch1 := make(chan int)
	go sendOnly(ch1, 10)
	go receiveOnly(ch1)
	time.Sleep(2 * time.Second)

	ch2 := sum(100)
	// fmt.Println(<-ch2)

	output := total(ch2) // output과 ch2는 수신전용 채널이다.
	// output <- 444 -> 그래서 이건 에러이다.
	fmt.Println(<-output)

}

// 발신 전용
func sendOnly(c chan<- int, cnt int) { // c chan<- int 이건 발신전용이기 때문에
	for i := 0; i < cnt; i++ {
		c <- i
	}
	c <- 777
	// <-c  이건 수신 코드라서 에러가 뜬다.
}

// 수신 전용
func receiveOnly(c <-chan int) {
	for val := range c {
		fmt.Println("received : ", val)
	}
	fmt.Println(<-c)
}

func sum(cnt int) <-chan int { // 수신전용 채널을 리턴타입으로
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 999
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for val := range c {
			fmt.Println("val : ", val)
			a += val
		}
		tot <- a
	}()
	return tot
}
