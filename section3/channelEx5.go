package section3

import (
	"fmt"
	"time"
)

func ChannelEx5() {
	// 채널 셀렉트 구문
	// 채널에 값이 수신되면 해당 case문 실행
	// 일회성 구문이므로, for문 안에서 수행한다.
	// default구문 처리 주의

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch2 <- "golang hi"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() { // 채널을 수발신하는 코드에선 default구문을 사용하면 안된다.
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
				// default:
				// 	fmt.Println("default test")
			}
		}
	}()

	time.Sleep(4 * time.Second) // 무한루프 for문은 main루틴이 끝나면 종료되므로 7초동안은 무한루프로 돌리겠다는 소리이다.

	ch3 := make(chan int)
	ch4 := make(chan string)
	go func() {
		for {
			num := <-ch3 // 발신
			fmt.Println("ch3 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch4 <- "golang hi" // 수신
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() { // 수발신 전부한다.
		for {
			select {
			case ch3 <- 777:
			case str := <-ch4:
				fmt.Println("ch4 : ", str)
				// default:
				// 	fmt.Println("default test")
			}
		}
	}()

	time.Sleep(4 * time.Second) // 무한루프 for문은 main루틴이 끝나면 종료되므로 7초동안은 무한루프로 돌리겠다는 소리이다.
}
