package cuboid

type Cuboid struct {
	L, B, H float32
}

func New(l, b, h float32) *Cuboid {
	return &Cuboid{L: l, B: b, H: h}
}
func (r *Cuboid) Area() float32 {
	return r.L * r.B * r.H
}

func (r *Cuboid) Perimeter() float32 {
	return 2 * (r.L + r.B + r.H)
}
