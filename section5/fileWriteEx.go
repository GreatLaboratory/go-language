package section5

import (
	"encoding/csv"
	"fmt"
	"os"
)

func FileWriteEx() {
	file1, err := os.Create("./files/test_write.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 파일 쓰기 1
	s1 := []byte{77, 78, 89, 56, 99}
	if f, err := file1.Write(s1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n", f)
	}
	file1.Sync() // 위 코드까지 파일에 쓰고 저장시킨 후에 다음 코드에서 또 쓰는 작업을 진행시킬 때 사용

	// 파일 쓰기 2
	s2 := "Hello Golang!!! \nI want you!! \nGood bye~~!\n"
	if f, err := file1.WriteString(s2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("쓰기 작업(2) 완료 (%d bytes) \n", f)
	}
	file1.Sync()

	// 파일 쓰기 3
	s3 := "WriteAt test"
	if f, err := file1.WriteAt([]byte(s3), 70); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n", f)
	}
	file1.Sync()

	// csv파일 생성
	file2, err2 := os.Create("./files/test_write.csv")
	if err2 != nil {
		fmt.Println(err)
	}

	wr := csv.NewWriter(file2)
	wr.Write([]string{"Kim", "4.5"})
	wr.Write([]string{"Lee", "4.1"})
	wr.Write([]string{"Park", "4.2"})
	wr.Write([]string{"Cho", "4.3"})
	wr.Write([]string{"Hong", "4.4"})
	wr.Write([]string{"Pyo", "4.5"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err3 := file2.Stat()
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Printf("csv 쓰기 작업 후 파일 크기 (%d bytes)\n", fi.Size())
	fmt.Printf("csv 파일명 (%s)\n", fi.Name())
	fmt.Println("운영체제 파일 권한 ", fi.Mode())

	// 리소스 해제
	defer file1.Close()
	defer file2.Close()
}
