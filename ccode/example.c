// example.c
#include <stdio.h>

void say_hello(const char* name) {
    printf("Hello, %s!\n", name);
}

void print_pointer(void* ptr) {
    printf("Pointer address: %p\n", ptr);
}

int add(int a, int b) {
    return a + b;
}

//clang -shared -o libexample.so -fPIC example.c  # Linux
//clang -shared -o libexample.dylib -fPIC example.c  # macOS
//clang -shared -o libexample.dll -fPIC example.c  # Windows

// what is -fPIC? Position Independent Code what does it mean? 
// It means that the generated machine code is not dependent on being located at a specific memory address.
// This is necessary for shared libraries, as they can be loaded at any memory address.
// how to create a non shared library?
// clang -o libexample.so example.c
// can I use this as well?
