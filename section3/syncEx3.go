package section3

import (
	"fmt"
	"runtime"
	"time"
)

// 동기화 사용하지 않을 경우의 예제
// 쓰기, 읽기의 동작순서가 일정하지 않아 잘못된 오류를 반환할 가능성 증가
func SyncEx3() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	go func() {
		for i := 1; i <= 10; i++ {
			data++
			fmt.Println("write : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("--read1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("----read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(5 * time.Second)
}
