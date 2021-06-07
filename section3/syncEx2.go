package section3

import (
	"fmt"
	"runtime"
	"sync"
)

// SyncEx2 동기화를 사용한 예제
// Mutex : 여러 고루틴에서 작업하는 공유데이터 보호
// Mutex : 상호배제 -> 쓰레드(고루틴)들이 서로 running time에 영향을 주지않게 단독으로 실행되게끔 하는 기술
// sync.Mutex 선언 후 Lock, Unlock 사용
func SyncEx2() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := Count2{
		num: 0,
	}

	ch := make(chan bool)
	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment2()
			ch <- true
			runtime.Gosched()
		}()
	}
	for i := 1; i <= 10000; i++ {
		<-ch
	}
	c.result2() // 이제 항상 10000이라는 값이 유지됨. -> 동기화 성공!!
}

type Count2 struct {
	num   int
	mutex sync.Mutex
}

func (c *Count2) increment2() {
	// 공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num++
	// 공유데이터 수정 후 보호 해제 -> 만약 unlock하지않으면 다른 고루틴들이 실행되지 않는다.
	c.mutex.Unlock()
}
func (c *Count2) result2() {
	fmt.Println("result : ", c.num)
}
