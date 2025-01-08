package shape

import "fmt"

func What() {
	whatIsthis() // this can be called with in the package
}

func whatIsthis() {
	fmt.Println("This package support rect as of now")
}

type T struct {
	a1, a2 int
	A3     int
}
