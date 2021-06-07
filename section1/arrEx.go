package section1

import "fmt"

func ArrEx() {
	// 배열은 용량, 길이가 항상 같다.
	// 배열 vs 슬라이스
	// 길이 고정 vs 길이 가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달  -- 7을 저장 vs 7이 저장되어있는 메모리의 주소를 전달. 받은 쪽에서 값을 수정하면 주소가 가리키는 값이 변경된다.
	// 전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가
	// 대부분 슬라이스 사용한다. 배열은 진짜 거의 사용 안한다.

	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 길이

	// 배열
	var arr1 [5]int
	var arr2 [3]int = [3]int{1, 2, 3}
	var arr3 = [5]int{1, 2, 3}
	arr4 := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}

	fmt.Println(arr1) // [0 0 0 0 0]
	arr1[2] = 5
	fmt.Println(arr1)      // [0 0 5 0 0]
	fmt.Println(arr2)      // [1 2 3]
	fmt.Println(arr3)      // [1 2 3 0 0]
	fmt.Println(arr4)      // [[1 2 3] [1 2 3] [1 2 3]]
	fmt.Println(len(arr4)) // 3

	for _, num := range arr3 {
		fmt.Println(num)
	}

	// 값은 복사되지만 주소는 서로 다름. 그래서 값이 바뀌어도 참조하는 주소는 서로 다르기에 한쪽만 값이 바뀜
	arr5 := [4]int{1, 10, 100, 1000}
	var arr6 [4]int
	arr6 = arr5
	fmt.Printf("%p %v\n", &arr5, arr5) // 0xc0000044a0 [1 10 100 1000]
	fmt.Printf("%p %v\n", &arr6, arr6) // 0xc0000044c0 [1 10 100 1000]

	arr6[1] = 999
	fmt.Println(arr5) // [1 10 100 1000]
	fmt.Println(arr6) // [1 999 100 1000]

}
