package section1

import (
	"fmt"
	"math"
)

func NumericEx() {
	// 암묵적 형변환은 안된다. true는 0이나 ""이나 nil로 표현 안된다.
	var num1 byte = 32
	var num2 byte = 0110
	var num3 byte = 0x48
	fmt.Printf("%c %c %c\n", num1, num2, num3)
	fmt.Printf("%d %d %d\n", num1, num2, num3)
	fmt.Printf("%d %o %x\n", num1, num2, num3)

	var num4 byte = 72
	var num5 byte = 0110
	var num6 byte = 0x48
	fmt.Printf("%c %c %c\n", num4, num5, num6)
	fmt.Printf("%d %d %d\n", num4, num5, num6)
	fmt.Printf("%d %o %x\n", num4, num5, num6)

	// var float1 float32 = 0.14
	// var float2 float32 = .369
	// var float3 float32 = 442.022335
	var float4 float32 = 10.0
	fmt.Println(float4 - 0.1)          // 9.9
	fmt.Println(float32(float4 - 0.1)) // 9.9
	fmt.Println(float64(float4 - 0.1)) // 9.899999618530273

	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}
