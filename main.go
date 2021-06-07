package main

import (
	"fmt"

	_ "practice/lib" // 이런 식으로 일단은 안쓰지만 나중에 쓸 수도 있는 패키지는 빈 식별자로 해놔야 포맷팅이 안된다.
	_ "practice/section2"
	"practice/section6"

	// 이렇게 별칭을 줄 수도 있다.

	good "practice/section1"
)

func init() {
	// init()함수는 패키지 로드 시에 가장 먼저 호출 (단 한번)
	fmt.Println("init func start!")
}

func main() {
	fmt.Println("this is real main entry file!!")
	fmt.Println(good.IsLargerThan(15, 80))
	fmt.Println(good.IsSmallerThan(15, 80))
	fmt.Println("random number => ", good.GetRandomNum())
	fmt.Println("-------------------------------------------------------")

	// section1.NumericEx()
	// section1.StringEx()
	// section1.ArrEx()
	// section1.SliceEx()
	// section1.MapEx()
	// section1.PointerEx()
	// section1.FunctionEx1()
	// section1.FunctionEx2()
	// section1.DeferEx()
	// section2.OOPEX1()
	// section2.OOPEX2()
	// section2.OOPEX3()
	// section2.InterfaceEx1()
	// section2.InterfaceEx2()
	// section3.GoRoutineEx()
	// section3.ChannelEx1()
	// section3.ChannelEx2()
	// section3.ChannelEx3()
	// section3.ChannelEx4()
	// section3.ChannelEx5()
	// section3.SyncEx1()
	// section3.SyncEx2()
	// section3.SyncEx3()
	// section3.SyncEx4()
	// section3.SyncEx5()
	// section3.SyncAdvancedEx()
	// section4.ErrorEx1()
	// section4.ErrorEx2()
	// section4.ErrorEx3()
	// section5.FileWriteEx()
	// section5.FileReadEx()
	// section5.FileIoUtilEx()

	// num := oper.Number{
	// 	X: 40,
	// 	Y: 20,
	// }
	// fmt.Println("합 : ", num.Plus())
	// fmt.Println("차 : ", num.Minus())
	// fmt.Println("곱 : ", num.Multiple())
	// fmt.Println("몫 : ", num.Divide())
	// fmt.Println("제곱의 합 : ", num.SquarePlus())
	// fmt.Println("제곱의 차 : ", num.SquareMinus())
	// fmt.Println("-------------------------------------------------------")

	// section6.ReadXlsx()
	section6.Crawler()
}
