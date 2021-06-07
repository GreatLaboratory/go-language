package math

// SquarePlus -> 두 숫자의 합
func (n *Number) SquarePlus() int {
	return (n.X * n.X) + (n.Y * n.Y)
}

// SquareMinus -> 두 숫자의 차
func (n *Number) SquareMinus() int {
	return (n.X * n.X) - (n.Y * n.Y)
}
