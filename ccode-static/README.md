## generate static library

clang -c -o example.o example.c

ar rcs libexample.a example.o 


#cgo CFLAGS: -I. tells the compiler to look for header files in the current directory.
#cgo LDFLAGS: -L. -lexample -static specifies the linker flags:
-L. tells the linker to look in the current directory for libraries.
-lexample links with libexample.a (the static library).
-static ensures that the library is statically linked (though on macOS, static linking is sometimes handled differently).

go build -ldflags "-extldflags '-static'" -o main



set CGO_ENABLED=1
set CC=gcc
set CGO_LDFLAGS=-L. -lexample

CGO_LDFLAGS=-L. -lexample
