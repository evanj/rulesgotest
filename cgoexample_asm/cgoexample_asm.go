package main

// #cgo CFLAGS: -Wall -Wextra -Werror

/*
#include "cfunc.h"

extern int example_asm_func();
*/
import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	C.example_c_func()

	asmResult := C.example_asm_func()
	fmt.Printf("example_asm_func returned value = %d\n", asmResult)
	fmt.Printf("should be 42 on arm64; 44 on amd64; GOARCH=%s\n", runtime.GOARCH)
}
