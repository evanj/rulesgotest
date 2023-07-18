package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("hello GOARCH=%s GOOS=%s\n", runtime.GOARCH, runtime.GOOS)
}
