package main

import "unsafe"

func main() {
	arr := [5]int{10, 11, 12, 13, 14}
	ptr1 := &arr[0] //
	//ptr1 += 8
	uintptr1 := uintptr(unsafe.Pointer(ptr1))
	for i := 0; i < len(arr); i++ {
		v := (*int)(unsafe.Pointer(uintptr1))
		//	A pointer value of any type can be converted to a Pointer.
		// A Pointer can be converted to a pointer value of any type.
		// A uintptr can be converted to a Pointer.
		// A Pointer can be converted to a uintptr.
		println(*v)
		uintptr1 += 8
	}
}
