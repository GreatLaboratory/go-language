package section5

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func FileReadEx() {
	if file, err := os.Open("./files/sample.txt"); err != nil {
		fmt.Println(err)
	} else {
		if fileInfo, err := file.Stat(); err != nil {
			fmt.Println(err)
		} else {
			fd1 := make([]byte, fileInfo.Size()) // 파일 크기 만큼의 슬라이스 만들기
			if ct1, err := file.Read(fd1); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("파일 정보 출력 : ", fileInfo)
				fmt.Println("파일 이름 : ", fileInfo.Name())
				fmt.Println("파일 사이즈 : ", fileInfo.Size())
				fmt.Println("파일 수정시간 : ", fileInfo.ModTime())
				fmt.Println("읽기 작업 완료 : ", ct1, " bytes")
				fmt.Println(string(fd1)) // 이게 파일 내용을 읽는다.
			}
		}
		if ol, err := file.Seek(20, 0); err != nil { // 두번째 인자 : 0이면 처음위치부터, 1이면 현재위치부터, 2는 마지막 위치
			fmt.Println(err)
		} else {
			fd2 := make([]byte, 20)
			if ct2, err := file.Read(fd2); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("읽기 작업 완료 : (%d bytes) (%d ret)\n", ct2, ol)
				fmt.Println(string(fd2)) // 이게 파일 내용을 읽는다.
			}
		}
		defer file.Close()
	}

	if file, err := os.Open("./files/sample.csv"); err != nil {
		fmt.Println(err)
	} else {
		rr := csv.NewReader(bufio.NewReader(file))
		if row, err := rr.Read(); err != nil { // 1개의 row 단위로 읽기
			fmt.Println(err)
		} else {
			for _, item := range row {
				fmt.Println("item : ", item)
			}
		}
		fmt.Println("============================================")
		if row, err := rr.Read(); err != nil { // 1개의 row 단위로 읽기
			fmt.Println(err)
		} else {
			for _, item := range row {
				fmt.Println("item2 : ", item)
			}
		}
		fmt.Println("============================================")
		if rows, err := rr.ReadAll(); err != nil { // 1개의 row 단위로 읽기
			fmt.Println(err)
		} else {
			for _, items := range rows {
				for _, item := range items {
					fmt.Println("item3 : ", item)
				}
			}
		}
		defer file.Close()
	}
}
