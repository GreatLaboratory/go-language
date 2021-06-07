package section3

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func SyncEx5() {
	// 고루틴 동기화 객체
	// 동기화 상태조건 메소드 사용
	// Wait, Notify, NotifyAll : 기타 언어
	// Wait, Signal, Broadcas : go언어

	runtime.GOMAXPROCS(runtime.NumCPU())

	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	ch := make(chan int, 5) // 뒤에 버퍼값을 주면 비동기다. - 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			ch <- 777 // 비동기 채널에게 777을 보낸다.
			fmt.Println("Go Routine Waiting : ", n)
			condition.Wait() // 고루틴이 대기, 멈춤 - 고루틴 재웠음
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("receive : ", <-ch)
	}

	// 한 개씩 깨운다. (모든 고루틴 생성 후)
	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Wake Go Routine(Signal) : ", i)
	// 	condition.Signal()
	// 	mutex.Unlock()
	// }

	// 한 번에 다 깨우기
	mutex.Lock()
	fmt.Println("Wake Go Routine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
