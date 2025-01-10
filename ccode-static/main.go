package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -Llib/. -lexample
#include "example.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Call C function `say_hello`
	name := C.CString("World")
	C.say_hello(name)

	// Call C function `add`
	result := C.add(3, 4)
	fmt.Printf("3 + 4 = %d\n", result)

	// Allocate a Go object
	goValue := 42

	// Get a pointer to the Go value
	goPointer := unsafe.Pointer(&goValue)

	// Pass the pointer to the C function
	C.print_pointer(goPointer)

	// Back in Go, print the pointer address for verification
	fmt.Printf("Go pointer address: %p\n", goPointer)
}

/*
Explanation of Each Part
	1.	-L.
	•	This tells the linker where to look for libraries during the linking phase.
	•	-L specifies a directory path for library files.
	•	. refers to the current directory.
	•	In this case, the linker is instructed to search for libraries in the current working directory where your Go code resides.
	2.	-lexample
	•	This specifies the name of the library to link against, without the lib prefix or the file extension.
	•	-l tells the linker to link against a library named example.
	•	For example:
	•	libexample.so on Linux.
	•	libexample.dylib on macOS.
	•	example.lib on Windows.
*/
