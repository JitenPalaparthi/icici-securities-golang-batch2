package main

func main() {

	// for
	// for range

	for i := 1; i <= 10; i++ {
		println(i)
	}

	num := 1

	for num <= 10 { // like a while loop
		println(num)
		num++
	}

	num = 1
	for {
		println(num)
		if num > 10 {
			break
		}
		num++
	}
	println("------")

	// done := false
	// for i := 1; i <= 10; i++ {
	// 	if done {
	// 		break
	// 	}
	// 	for j := 1; j <= 10; j++ {
	// 		if i == j {
	// 			done = true
	// 			break
	// 		}
	// 		println(i, "->", j)
	// 	}
	// }

out:
	for i := 4; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == j {
				break out
			}
			println(i, "->", j)
		}
	}

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue // continues to the next iteration
		}
		println(i)
	}

	i := 1

	for ; ; i++ {
		println(i)
	}

}
