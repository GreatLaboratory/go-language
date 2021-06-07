package section1

import (
	"fmt"
	"os"
)

func packageEx() {
	// 패키지 이름 == 디렉토리 이름
	// 같은 패키지 내의 소스파일들은 디렉토리명을 패키지 이름으로 사용한다.
	// main 패키지는 특별하게 인식된다.
	// 네이밍 규칙 -> 소문자 private / 대문자는 public

	var name string
	fmt.Println("이름은 ? ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}

func IsLargerThan(x int32, y int32) bool {
	return x > y
}
