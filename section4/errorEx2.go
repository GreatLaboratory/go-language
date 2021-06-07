package section4

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"
)

func ErrorEx2() {
	if v, err := power2(0, 5); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(v)
	}
	fmt.Println("hi") // 위에서 에러가 걸리면 아래는 실행안됨
}

// 에러 처리 구조체
type PowError struct {
	time    time.Time "에러 발생 시간"
	value   float64   "파라미터"
	message string    "에러 메세지"
}

func (e *PowError) Error() string {
	// err := errors.New("[]")
	// return err.Error()
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func power2(f, i float64) (float64, error) {
	if f == 0 {
		powErr := PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
		return 0, errors.New(powErr.Error())
	}
	if math.IsNaN(f) {
		powErr := PowError{time: time.Now(), value: f, message: "숫자를 사용해야 합니다."}
		return 0, errors.New(powErr.Error())
	}
	if math.IsNaN(i) {
		powErr := PowError{time: time.Now(), value: f, message: "숫자를 사용해야 합니다."}
		return 0, errors.New(powErr.Error())
	}
	return math.Pow(f, i), nil
}
