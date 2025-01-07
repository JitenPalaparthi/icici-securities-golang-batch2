package main

func main() {

	day := 9

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("No day")
	}

	// empty switch

	num := -100

	switch { // empty switch
	case num > 0 && num <= 50:
		println(num, " is between 0 and 50")
	case num > 50 && num <= 100:
		println(num, " is above 50 and less than equal to 100")
	default:
		println(num, "num is either negative or above 100")
	}

	//

	char := 'a'

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U': // case can be any one of these
		println(string(char), " is vovel")
	default:
		println(string(char), " is either a consonent or a special unicode chars")
	}

	char = 'a'
	switch {
	case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U': // case can be any one of these
		println(string(char), " is vovel")
	default:
		println(string(char), " is either a consonent or a special unicode chars")
	}

	num = 48
	switch {
	case num%8 == 0:
		println(num, " is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, " is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, " is divisible by 2")
	}

	num = 42
	switch {
	case num%4 == 0:
		println(num, " is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, " is divisible by 2")
		fallthrough
	case num%8 == 0:
		println(num, " is divisible by 8")
	}

}
