package section3

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func GoRoutineEx() {
	// 타언어의 쓰레드와 비슷한 기능
	// 생성방법 매우 간단, 리소스 매우 적게 사용
	// 비동기적 함수 루틴 실행 -> 채널을 통한 통신 가능
	// 공유메모리 사용 시에 정확한 동기화 코딩 필요

	// exe1()
	// fmt.Println("Main Routine Start", time.Now())
	// go exe2() // 메인 루틴이 종료가 되어서 해당 서브 루틴이 실행안된다. // 메인이 끝나면 자식은 끝났든 말든 다 종료시킨다.
	// go exe3() // exe2랑 exe3은 서로 비동기적으로 경쟁상태에 들어가서 순서가 보장되지 않는다.
	// time.Sleep(4 * time.Second)
	// fmt.Println("Main Routine End", time.Now())

	//------------------------------------------------------

	// exe4("t1")
	// fmt.Println("Main Routine Start : ", time.Now())
	// go exe4("t2") // 각각 1000번 반복 for문인데 서로 시간순서상 섞여서 실행된다.
	// go exe4("t3")
	// time.Sleep(4 * time.Second)
	// fmt.Println("Main Routine End : ", time.Now())

	//------------------------------------------------------

	// 멀티 코어(다중cpu) 최대한 활용
	// cpuNum := runtime.GOMAXPROCS(runtime.NumCPU()) // 현 시스템의 cpu코어 개수 반환 후 설정
	// fmt.Println("Current System : CPU : ", cpuNum) // 설정값 출력
	// fmt.Println("Main Routine Start : ", time.Now())
	// for i := 0; i < 10; i++ {
	// 	go exe5(i)
	// }
	// time.Sleep(5 * time.Second)
	// fmt.Println("Main Routine End : ", time.Now())

	//------------------------------------------------------
	// 클로저 사용 예제 - 많이 사용하진 않음
	runtime.GOMAXPROCS(1)
	s := "GoRoutine Closure : "
	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i + 1) // 반복문 클로저는 일반적으로 즉시 실행인데 고루틴 클로저는 가장 나중에 반복문이 끝난 후 실행
	}
	time.Sleep(4 * time.Second)

}

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초라는 뜻
	fmt.Println("exe1 func end", time.Now())
}
func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초라는 뜻
	fmt.Println("exe2 func end", time.Now())
}
func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초라는 뜻
	fmt.Println("exe3 func end", time.Now())
}

func exe4(name string) {
	fmt.Println(name, "start : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>", i+1)
	}
	fmt.Println(name, "end : ", time.Now())
}

func exe5(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}
