package section3

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// SyncEx4 동기화를 사용할 경우의 예제
// RWMutex : 쓰기 Lock -> 쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안된다. => 읽기 락, 쓰기 락 전부 방지한다.
// RMutex : 읽기 Lock -> 읽기 시도 중에 값 변경 방지 => 쓰기 락만 방지한다.
func SyncEx4() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()

			data++
			fmt.Println("write : ", data)
			time.Sleep(200 * time.Millisecond)

			// 쓰기 뮤텍스 잠금 해제
			mutex.Unlock()
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()

			fmt.Println("--read1 : ", data)
			time.Sleep(1 * time.Second)

			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()

			fmt.Println("----read2 : ", data)
			time.Sleep(1 * time.Second)

			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()
	time.Sleep(10 * time.Second)
}
