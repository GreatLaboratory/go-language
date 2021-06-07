package math

// Number -> 두 개의 숫자 구조체
type Number struct {
	X int
	Y int
}

// Plus -> 두 숫자의 합
func (n *Number) Plus() int {
	return n.X + n.Y
}

// Minus -> 두 숫자의 차
func (n *Number) Minus() int {
	return n.X - n.Y
}

// Multiple -> 두 숫자의 곱
func (n *Number) Multiple() int {
	return n.X * n.Y
}

// Divide -> 두 숫자를 나눴을 때의 몫
func (n *Number) Divide() int {
	return n.X / n.Y
}
