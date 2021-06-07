package section1

import "fmt"

func MapEx() {
	// key-value
	// 참조 타입 (참조 값 전달)
	// 그래서 비교 연산자가 당연히 사용 불가능
	// key는 참조값이 안되고 value는 모든 타입 사용가능
	// make() 및 리터럴로 초기화 가능
	// 순서가 없음

	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)
	map3 := make(map[string]int)
	fmt.Println(map1) // map[]
	fmt.Println(map2) // map[]
	fmt.Println(map3) // map[]

	map4 := map[string]int{}
	map4["aa"] = 12
	map4["bb"] = 22
	map4["cc"] = 32

	map5 := map[string]int{
		"aa": 456,
		"bb": 123,
		"cc": 789,
	}

	map6 := make(map[string]int, 10)
	map6["aa"] = 12
	map6["bb"] = 22
	map6["cc"] = 32

	fmt.Println(map4)       // map[aa:12 bb:22 cc:32]
	fmt.Println(map5)       // map[aa:456 bb:123 cc:789]
	fmt.Println(map6)       // map[aa:12 bb:22 cc:32]
	fmt.Println(map6["aa"]) // 12
	fmt.Println(map6["cc"]) // 32

	// 조회 및 순회
	map7 := map[string]string{
		"daum":   "daum.net",
		"naver":  "naver.com",
		"nate":   "nate.com",
		"google": "google.com",
	}
	for key, value := range map7 { // 순서 없으므로 랜덤
		fmt.Println(key)
		fmt.Println(value)
	}

	// 값 추가 및 변경, 삭제
	map7["newbie"] = "newbie"
	fmt.Println(map7) // map[daum:daum.net google:google.com nate:nate.com naver:naver.com newbie:newbie]
	map7["newbie"] = "updatedNewbie"
	fmt.Println(map7) // map[daum:daum.net google:google.com nate:nate.com naver:naver.com newbie:updatedNewbie]
	delete(map7, "daum")
	delete(map7, "naver")
	fmt.Println(map7) // map[google:google.com nate:nate.com newbie:updatedNewbie]

	// 존재하지 않는 키에 대한 핸들링
	fruits := map[string]int{
		"apple":  1,
		"banana": 2,
		"lemon":  0,
	}
	value1, isExist1 := fruits["lemon"]
	value2 := fruits["kiwi"]           // 존재하지않는 키값은 value값타입의 초기값으로 설정
	value3, isExist2 := fruits["kiwi"] // 근데 이러면 lemon의 값이 0인것처럼 없는 키일 수도 있으니까 헷갈려서 인자를 하나 더 받는다. 키 존재여부
	fmt.Println(value1, isExist1)      // 0 true
	fmt.Println(value2)                // 0
	fmt.Println(value3, isExist2)      // 0 false

	if value, isExist := fruits["kiwi"]; isExist {
		fmt.Println(value)
	} else {
		fmt.Println("there is no kiwi")
	}

	if _, isExist := fruits["kiwi"]; !isExist {
		fmt.Println("there is no kiwi")
	}
}
