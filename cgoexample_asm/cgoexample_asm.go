package main

// #cgo CFLAGS: -Wall -Wextra -Werror

/*
#include "cfunc.h"

int example_asm_func();
*/
import "C"
import "fmt"

func main() {
	C.example_c_func()

	asmResult := C.example_asm_func()
	fmt.Printf("example_asm_func returned value = %d\n", asmResult)
}
