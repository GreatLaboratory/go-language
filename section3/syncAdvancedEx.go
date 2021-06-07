package section3

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func SyncAdvancedEx() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Once : 한번만 실행(주로 초기화에 사용)
	// Do로 실행한다.
	once := new(sync.Once)
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Go Routine : ", n)
			once.Do(onceTest) // 5번 실행되는게 아니라 딱 한번만 실행되게끔 한다.
		}(i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("---------------------------------")

	// 대기그룹 : 실행 흐름 변경 (고루틴이 종료될 때까지 대기 기능)
	// Wait Group : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시까지 대기)
	// Add로 추가된 고루틴 개수와 Done으로 종료된 알림횟수가 같아야 정확하게 동작한다.

	// 원자성 사용 안한 경우
	waitGroup := new(sync.WaitGroup)
	cnt := 0
	for i := 0; i < 5000; i++ {
		waitGroup.Add(1)
		go func(n int) {
			// fmt.Println("WaitGroup : ", n)
			cnt++
			waitGroup.Done()
		}(i)
	}
	for i := 0; i < 2000; i++ {
		waitGroup.Add(1)
		go func(n int) {
			// fmt.Println("WaitGroup : ", n)
			cnt--
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()                           // 대기그룹이 끝날 때까지 기다린다. 100번이 add됐으면 100번 done될 때까지 메인스레드가 종료되지않고 기다린다.
	fmt.Println("Wait Group End? >>>>> ", cnt) // 원자성이 보장되려면 3000이 나와야하는데 그렇지 못한다.

	fmt.Println("------------------------------------------------")

	// 원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작 (트랜젝션이랑 비슷함)
	// 모든 조작이 완료될 때까지 다른 프로세스 개입 불가
	// 주로 공용 변수에 관한 계산에서 사용

	// 원자성 사용한 경우
	waitGroup2 := new(sync.WaitGroup)
	var cnt2 int64 = 0
	for i := 0; i < 5000; i++ {
		waitGroup2.Add(1)
		go func(n int) {
			// fmt.Println("WaitGroup : ", n)
			// cnt2++
			atomic.AddInt64(&cnt2, 1)
			waitGroup2.Done()
		}(i)
	}
	for i := 0; i < 2000; i++ {
		waitGroup2.Add(1)
		go func(n int) {
			// fmt.Println("WaitGroup : ", n)
			// cnt2--
			atomic.AddInt64(&cnt2, -1)
			waitGroup2.Done()
		}(i)
	}
	waitGroup2.Wait()
	finalCnt := atomic.LoadInt64(&cnt2)
	fmt.Println("Wait Group End? >>>>> ", cnt2) // 원자성이 보장되어서 항상 3000이 나온다.
	fmt.Println("Wait Group End? >>>>> ", finalCnt)
}
func onceTest() {
	fmt.Println("Once Execute only")
}
