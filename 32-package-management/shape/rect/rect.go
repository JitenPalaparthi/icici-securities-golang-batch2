package rect

type Rect struct {
	L, B float32
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func NewDefault() *Rect {
	return &Rect{L: 1, B: 1}
}
