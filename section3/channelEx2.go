package section3

import "fmt"

func ChannelEx2() {
	ch1 := make(chan int)

	go rangeSum(501, ch1) // 여기서 각 쓰레드에서 채널을 통해 값을 ch1로 보내면
	go rangeSum(502, ch1)
	go rangeSum(500, ch1)

	// 여기 메인 쓰레드에서 그 ch1값을 받는다.
	// 채널을 통해 값을 수신받을 때까지 기다린다.
	// 순서대로 데이터 수신(동기) : 채널에서 값 수신 완료될 때까지 대기
	result1 := <-ch1
	result2 := <-ch1
	result3 := <-ch1
	fmt.Println(result1, result2, result3)

	fmt.Println("------------------------------")

	ch2 := make(chan bool)

	go func() {
		for i := 1; i <= 5; i++ {
			ch2 <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 1; i <= 5; i++ {
		<-ch2 // 위에 go 루틴 쓰레드가 여기서 채널을 통해 수신받을 때까지 대기한다.
		fmt.Println("Main : ", i)
	}

	fmt.Println("------------------------------")

	// 버퍼 : 발신할 땐 가득차면 대기 - 비어있으면 작동
	// 수신할 땐 비어있으면 대기하고 가득차면 작동
	// 이건 비동기 흐름이다.
	ch3 := make(chan bool, 2) // 여기서 두번째 인자가 버퍼이다.

	go func() {
		for i := 0; i < 12; i++ {
			ch3 <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < 12; i++ {
		<-ch3 // 위에 go 루틴 쓰레드가 여기서 채널을 통해 수신받을 때까지 대기한다.
		fmt.Println("Main : ", i)
	}
}

func rangeSum(n int, c chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	c <- sum
}
