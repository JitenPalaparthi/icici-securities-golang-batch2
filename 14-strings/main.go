package main

func main() {

	str1 := "Hello World!" // collection of bytes

	str2 := "Hello 世界!"
	var str3 string

	print(str3)

	// str1 header
	// ptr =
	// len = 12

	// str3 header
	// ptr: 0x11 // it is not nil , it has pointed to some address, some address which is uused just as a placeholder or dummy
	// len: 0
	println("len of str1", len(str1))
	println("len of str2", len(str2))

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]))
	}

	println()
	for i := 0; i < len(str2); i++ {
		print(string(str2[i]))
	}
	println()

	for i, v := range str2 {
		println(i, "-->", string(v))
	}
	// can nil be checked on string?

}

// uni8
// 0-255 - 1
// 255- 16365 - 2
// 16360- 32,000 -3
// 32,000 - 4
