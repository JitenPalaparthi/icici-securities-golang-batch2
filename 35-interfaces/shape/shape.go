package shape

type IShape interface {
	Area() float32      // signature of the method
	Perimeter() float32 // signature of the method
	IWhat
}

type IWhat interface {
	What() string
}
