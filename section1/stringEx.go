package section1

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringEx() {

	var str1 string = "hello"
	var str2 string = "한글"

	fmt.Println(len(str1))                    // 영어는 제대로 나오는데
	fmt.Println(len(str2))                    // 한국어는 제대로 안나온다
	fmt.Println(len([]rune(str2)))            // 제대로 나오게 하는 방법
	fmt.Println(utf8.RuneCountInString(str2)) // 제대로 나오게 하는 방법

	fmt.Println(str2[0])                                    // 이러면 아스키코드값이 나온다. 237
	fmt.Printf("%c %c\n", str2[0], str2[1])                 // 문자열이 나온다. 근데 한글 깨짐
	fmt.Printf("%c %c\n", []rune(str2)[0], []rune(str2)[1]) // 문자열이 나온다. 한 글 -> 한글 안깨짐

	for i, char := range str1 {
		fmt.Printf("%d - %c\t", i, char)
	}
	fmt.Println()
	for j, char := range []rune(str2) {
		fmt.Printf("%d - %c\t", j, char)
	}
	fmt.Println()

	var str3 string = "Golang"
	var str4 string = "World"

	fmt.Println(str3[0:2], str3[0]) // go 71
	fmt.Println(str4[3:], str4[0])  // ld 87
	fmt.Println(str4[:4])           // worl
	fmt.Println(string(str4[0]))    // w

	fmt.Println(str3 == str4) // flase
	fmt.Println(str3 != str4) // true
	fmt.Println(str3 > str4)  // flase - 문자열 비교할 때 아스키 코드에 대한 사전식 비교를 한다.
	fmt.Println(str3 < str4)  // true

	str5 := "hi " + "golang " + "world "
	str6 := "good"
	fmt.Println(str5 + str6) // hi golang world good

	// append, join작업이 더 효율이 좋다. -------------->>>>> 중요하다.
	strSet := []string{} //슬라이스 선언
	strSet = append(strSet, str5)
	strSet = append(strSet, str6)
	fmt.Println(strings.Join(strSet, "---")) // hi golang world  ---good

}
