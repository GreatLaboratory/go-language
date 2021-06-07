package section3

import "fmt"

func ChannelEx3() {
	// close : 채널 닫기 - 닫힌 채널에 값 전송 시 패닉 발생
	// range : 채널안에서 값을 순회하여 꺼낸다. 채널을 닫아야 반복문이 종료, 채널이 열려있고 값 전송하지않으면 계속 대기하는 문제발생
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1) // 5회 채널에 값 전송 후 채널 닫기
	}()

	// 채널이 close될 때까지 채널에서 값을 꺼내온다.
	for v := range ch1 { // 채널의 닫은 상태를 확인할 때까지 채널에서 값이 도착하기를 대기하는거임
		fmt.Println(v) // 0 1 2 3 4 순서가 보장된다.
	}

	fmt.Println("----------------------------------------")

	ch2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- "Golang"
		}
	}()

	if v1, ok := <-ch2; ok {
		fmt.Println(v1, ok)
	}
	if v2, ok := <-ch2; ok {
		fmt.Println(v2, ok)
	}
	if v3, ok := <-ch2; ok {
		fmt.Println(v3, ok)
	}
	close(ch2)
	if v4, ok := <-ch2; !ok {
		fmt.Println(v4, ok) // " " false
	}
}
