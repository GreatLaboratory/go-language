package section5

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func FileIoUtilEx() {
	// ioutil 사용
	str := "Hello Golang\nFile Write Test\n"
	err := ioutil.WriteFile("./files/test_write1.txt", []byte(str), os.FileMode(0644))

	data, err := ioutil.ReadFile("./files/sample.txt")
	errCheck(err)
	fmt.Println(string(data))

	// bufio 사용
	file, err := os.OpenFile("./files/test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777)) // 없으면 만들고 있으면 읽고
	errCheck(err)

	// wt := bufio.NewWriter(file) // Writer 반환 (버퍼 사용)
	// wt.WriteString("Hello Golang!\nFile Write Test1\n")
	// wt.Write([]byte("Hello Golang!\nFile Write Test2\n"))
	// fmt.Printf("사용한 Buffer size (%d bytes)\n", wt.Buffered())
	// fmt.Printf("사용한 Buffer size (%d bytes)\n", wt.Available())
	// fmt.Printf("사용한 Buffer size (%d bytes)\n", wt.Size())
	// wt.Flush() // 이게 없으면 버퍼에만 값이 차있고 파일은 써지지않음 -> 버퍼를 비우고 디스크에 기록시킨다.

	rt := bufio.NewReader(file) // Reader 반환
	fi, err := file.Stat()
	errCheck(err)
	b := make([]byte, fi.Size())
	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())
	rt.Read(b)
	fmt.Println(string(b))

	defer file.Close()
}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}
