package main

import "fmt"

func main() {
	var n4 int = 12321 //00000000 00000000 00000000 00000000 00000000 00000000 00110000 00100001
	fmt.Printf("%b\n", n4)
	var n5 uint16 = uint16(n4)
	fmt.Println(n5)

	arr := [2]int{10, 20}
	slice := arr[:]
	SumOf(slice...)
	SumOf(arr[:]...)

	c := 100*200/1 + 23 + n5 + uint16(n4)

	i := 0
loop: //label
	//println(i)
	i++
	if i > 10 {
		goto exit
	}
	if i%2 == 0 {
		goto even
	} else {
		goto odd
	}

even:
	println("Even:", i)
	if i <= 10 {
		goto loop
	}
odd:
	println("Odd:", i)
	if i <= 10 {
		goto loop
	}
exit:
	println("finished")
}

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
