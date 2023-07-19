package main

/*
#include <stdio.h>
#include <resolv.h>

static void example_c_func() {
	printf("example_c_func! func addr=%p\n", example_c_func);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.example_c_func()

	// from src/net/cgo_unix_cgo_darwin.go
	fmt.Printf("sizeof(struct __res_state)=%d\n", unsafe.Sizeof(C.struct___res_state{}))
}
