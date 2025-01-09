package main

func main() {

	slice := []int{1, 2, 3, 4, 5}

	fnslice1 := make([]func(), 0)
	fnslice2 := make([]func(), 0)
	fnslice3 := make([]func(), 0) // solve the problem of fnslice2

	//for _, v := range slice { // a problem was solved with loop in 1.22 onwards
	for i := 1; i < len(slice); i++ {
		fnslice1 = append(fnslice1, func() {
			println(i)
		})
	}

	i := 0
loop:
	fnslice2 = append(fnslice2, func() {
		println(i)
	})
	i++
	if i < len(slice) {
		goto loop
	}

	for _, fn := range fnslice1 {
		fn()
	}

	for _, fn := range fnslice2 {
		fn()
	}

	j := 0
loop1:
	j1 := j // this is what exactly a new loop is doing
	fnslice3 = append(fnslice3, func() {
		println(j1)
	})
	j++
	if j < len(slice) {
		goto loop1
	}

	for _, fn := range fnslice3 {
		fn()
	}

}
