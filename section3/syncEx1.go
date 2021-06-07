package section3

import (
	"fmt"
	"runtime"
)

// SyncEx1 -> 동기화를 사용하지 않은 예제 (비동기)
func SyncEx1() {
	// 실행 흐름 제어 및 변수 동기화 가능
	// 공유 데이터 보호가 가장중요

	// 아래는 동기화를 사용하지 않은 예제
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := Count{
		num: 0,
	}

	ch := make(chan bool)
	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment() // num값을 증가시킬 때마다 밑에 있는걸로 채널에 true를 보내준다.
			ch <- true
			runtime.Gosched() // 다른 cpu에게 양보해주는 함수
		}()
	}
	for i := 1; i <= 10000; i++ {
		<-ch // 이렇게 밑에서 수신을 해야 동기화 처리가 된다.
	}
	c.result() // 값이 항상 보장되지 않는다. -> 계속 다른 값이 출력된다. 이게 바로 동기화
	// 여러개의 코어 cpu에서 경쟁상태로 c.name이라는 변수에 값을 부여하는 increment함수를 go routine에서 실행시키기 때문이다.
}

// Count 여러 go routine이 공유하는 데이터
type Count struct {
	num int
}

func (c *Count) increment() {
	c.num++
}
func (c *Count) result() { // 여기선 값을 변경해서 저장하는 목적이 아닌 오직 readonly이기 때문에 매개변수에 포인터가 없어도 된다.
	fmt.Println("result : ", c.num)
}
