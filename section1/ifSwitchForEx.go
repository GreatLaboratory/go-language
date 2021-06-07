package section1

import (
	"fmt"
	"math/rand"
	"time"
)

func IsSmallerThan(x int32, y int32) bool {
	return x < y
}

func GetRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func ifSwitchForEx() {
	// if문
	// 1 또는 0이 조건으로 될 수 없다. 아쉽네..
	if num3 := 44; num3 > 11 {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
	}
	// fmt.Println(num3) -> 에러

	// switch
	// break가 필요없다. 알아서 맞는거 찾으면 뒤에있는건 실행 안한다. 자동break  fallthrough
	// 값이 아닌 type으로 분기 가능하다.
	// switch 뒤에 표현식 생략 가능
	// case 뒤에 표현식 사용 가능
	switch aaa := 333 / 2; aaa {
	case 2, 4, 5, 166:
		fmt.Println("짝수", aaa)
	case 9, 6, 33:
		fmt.Println("홀수", aaa)
	}

	switch thisOne := "go"; thisOne + "lang" {
	case "python":
		fmt.Println("this is python")
	case "go":
		fmt.Println("this is go")
	case "golang":
		fmt.Println("this is golang! correct")
	case "java":
		fmt.Println("this is java")
	default:
		fmt.Println("no contents")
	}

	rand.Seed(time.Now().UnixNano()) // 이거때매 매번 컴파일 할때마다 다른 값이 나온다. 매번 시간은 달라지니까
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println(i, "50 이상 100 미만")
	case i >= 0 && i < 50:
		fmt.Println(i, "0 이상 50 미만")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("java")
	case "go":
		fmt.Println("go")
		fallthrough // 이거 덕분에 go에서 걸렸음에도 다음 case문을 무조건 실행시킨다. ㄷㄷ
	case "python":
		fmt.Println("python")
		fallthrough
	case "js":
		fmt.Println("js")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	loc := []string{"seoul", "busan", "hanam"}
	for _, name := range loc { // 첫번째 매개변수는 무조건 index이다. 근데 _ 사용하면 묵음.
		fmt.Println(name)
	}

	// 후치연산은 반환값이 없다. ex) j := i++ -> x
	// 전위연산은 아예 지원을 안한다. ex) ++i
	sum, i := 0, 0
	for i <= 100 {
		sum += i // 짧은 연산에서 초기화는 한번이지만 대입은 또 된다.
		i++
	}
	fmt.Println(sum)

	sum2, j := 0, 0
	for { // 무한루프
		if j > 100 {
			break
		}
		sum2 += j
		j++
	}
	fmt.Println(sum2)

	for n, m := 0, 0; n <= 10; n, m = n+1, m+10 {
		fmt.Println(n, m)
	}

Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1 // 이러면 Loop1와 가장 가까운 반복문을 탈출한다.
				// break  // 만약에 Loop1없이 그냥 break만 하면 가장 가까운 반복문인 j반복문만 탈출해서 i가 3인 경우부터 끝까지 다 출력된다.
			}
			fmt.Println(i, j)
		}
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("---------------------")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2 // 이러면 바로 가장 Loop3과 가까운 반복문을 시작한다. 아래에 있는 프린트문은 실행x
			}
			fmt.Println(i, j)
		}
	}

}
