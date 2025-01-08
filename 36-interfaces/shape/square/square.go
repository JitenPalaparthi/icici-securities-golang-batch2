package square

type Square float32

func New(s float32) Square {
	return Square(s)
}
func (r Square) Area() float32 {
	return float32(r * r)
}

func (r Square) Perimeter() float32 {
	return 4 * float32(r)
}

func (r Square) What() string {
	return "Square"
}
