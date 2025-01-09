package main

import "errors"

func main() {
	//panic("Something went wrong")
	// if true {
	// 	fmt.Println("True")
	// 	return
	// }
	//Fatal("somethign went wrong", "Do you know what is that", true)

	fn2 := FullNameStruct{fn: "Jiten", ln: ""}
	r2 := fn2.GetFullName()
	println(r2)
	fn1 := GetFullName("Jiten", "P")
	println(fn1)

}

// func Fatal(args ...any) {
// 	fmt.Println(args)
// 	os.Exit(1) // non zero is crashed exit
// }

//

func GetFullName(fn, ln string) string {
	if fn == "" {
		panic("firstname cannot be empty")
	}
	if ln == "" {
		panic("lastname cannot be empty")
	}
	return fn + ln
}

func GetFullName1(fn, ln string) (string, error) {
	if fn == "" {
		return "", errors.New("firstname cannot be empty")
	}
	if ln == "" {
		return "", errors.New("lastname cannot be empty")
	}
	return fn + ln, nil
}

type FullNameStruct struct {
	fn, ln string
}

func (f *FullNameStruct) GetFullName() string {
	if f.fn == "" {
		panic("firstname cannot be empty")
	}
	if f.ln == "" {
		panic("lastname cannot be empty")
	}
	return f.fn + f.ln
}
