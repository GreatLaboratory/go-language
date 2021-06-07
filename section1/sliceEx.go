package section1

import (
	"fmt"
	"sort"
)

func SliceEx() {
	// 배열 vs 슬라이스
	// 길이 고정 vs 길이 가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달  -- 7을 저장 vs 7이 저장되어있는 메모리의 주소를 전달. 받은 쪽에서 값을 수정하면 주소가 가리키는 값이 변경된다.
	// 전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가
	// 대부분 슬라이스 사용한다.
	// make()함수로 선언했을 때 3번째 인자가 생략되면 용량과 길이가 같아진다.
	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 길이

	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4}
	slice4 := [][]int{{1, 2}, {3, 4}}
	slice3[2] = 10
	// slice2[0] = 1 // panic: runtime error: index out of range [0] with length 0
	fmt.Println(slice1, len(slice1), cap(slice1)) // [] 0 0
	fmt.Println(slice2, len(slice2), cap(slice2)) // [] 0 0
	fmt.Println(slice3, len(slice3), cap(slice3)) // [1 2 10 4] 4 4
	fmt.Println(slice4, len(slice4), cap(slice4)) // [[1 2] [3 4]] 2 2
	if slice1 == nil {
		fmt.Println("this is nil slice")
	}

	// make() - 슬라이스 선언
	// 슬라이스의 길이는 5이고 용량은 10이라 메모리를 10만큼 확보해둔다. 용량이 초과되면 메모리 재할당을 한다.
	slice5 := make([]int, 5, 10)
	slice5[2] = 7
	fmt.Println(slice5, len(slice5), cap(slice5)) // [0 0 7 0 0] 5 10

	// 값도 복사되고 메모리 주소도 복사됨. 그래서 값이 바뀌어도 같은 주소를 참조하여 둘 다 값이 바뀜
	slice6 := []int{1, 5, 9}
	slice7 := make([]int, 3)
	slice7 = slice6
	slice7[1] = 777
	fmt.Printf("%p %v\n", &slice6, slice6) // 0xc0000044c0 [1 777 9] -> 배열에선 처음 초기화값이 그대로 남아있음
	fmt.Printf("%p %v\n", &slice7, slice7) // 0xc0000044a0 [1 777 9]
	for _, num := range slice7 {
		fmt.Println(num)
	}

	// 예외 상황
	slice8 := make([]int, 50, 100)
	fmt.Println(slice8[4]) // 0
	//fmt.Println(slice8[50]) // panic: runtime error: index out of range -> 용량이 아닌 길이

	// 원소 추가
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10}
	s3 := []int{11, 12, 13}
	s1 = append(s1, 6, 7) // 길이값과 용량값이 2씩 증가된다.
	// s1의 주소값을 다시 s1에 재할당
	// 근데 용량을 만약 10으로 잡았다면 길이는 7이지만 용량은 아직 10이다.

	// 슬라이스 병합 -  ...을 붙여야한다.
	s2 = append(s1, s2...)
	s3 = append(s2, s3[0:2]...)
	fmt.Println(s1) // [1 2 3 4 5 6 7]
	fmt.Println(s2) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(s3) // [1 2 3 4 5 6 7 8 9 10 11 12]

	// 슬라이스의 용량이 넘칠 경우 2배만큼 재할당된다.
	s4 := make([]int, 3)
	s4 = append(s4, 9999)            // 용량과 길이가 3인 슬라이스에 1만큼 초과되면 용량은 2배로 재할당해준다.
	fmt.Println("length: ", len(s4)) // 4
	fmt.Println("weight: ", cap(s4)) // 6
	fmt.Println(s4)                  // [0 0 0 9999]
	s5 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s5 = append(s5, i)
		fmt.Println("length : ", len(s5), ", capacity : ", cap(s5))
	}

	// 추출
	s6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s6[1:3]) // 1번 인덱스부터 2번 인덱스까지
	fmt.Println(s6[:3])  // 처음부터 2번 인덱스까지
	fmt.Println(s6[3:])  // 3번 인덱스부터 끝까지
	fmt.Println(s6[:])   // 전체

	// 정렬
	s7 := []int{3, 6, 4, 2, 1, 98, 54, 22, 11, 66, 77, 84}
	s8 := []string{"a", "b", "t", "r", "w", "j", "l"}
	fmt.Println(sort.IntsAreSorted(s7))    // false
	fmt.Println(sort.StringsAreSorted(s8)) // false
	sort.Ints(s7)
	sort.Strings(s8)
	fmt.Println(s7)                        // [1 2 3 4 6 11 22 54 66 77 84 98]
	fmt.Println(s8)                        // [a b j l r t w]
	fmt.Println(sort.IntsAreSorted(s7))    // true
	fmt.Println(sort.StringsAreSorted(s8)) // true

	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make()으로 공간을 할당 후 복사해야함.
	// 복사된 슬라이스 값 변경해도 원본에는 영향 없음 -> 핵심
	ss1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss2 := make([]int, 5)
	ss3 := []int{}
	copy(ss2, ss1)
	copy(ss3, ss1)
	fmt.Println(ss2) // [1 2 3 4 5]
	fmt.Println(ss3) // [] -> make()로 공간 할당을 안했기때문에 복사가 안됨.
	ss2[1] = 9999
	fmt.Println(ss1) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(ss2) // [1 9999 3 4 5]

	ss4 := []int{1, 2, 3}
	ss5 := ss4[0:2] // 추출해도 똑같이 주소를 참조한다.
	ss5[0] = 888
	fmt.Println(ss4) // [888 2 3]
	fmt.Println(ss5) // [888 2]

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[0:5:7]                  // 0에서 4번 인덱스까지 추출하고 용량은 7로 할당
	fmt.Println(len(b), cap(b), b) // 5 7 [1 2 3 4 5]

}
