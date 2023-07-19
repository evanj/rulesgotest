package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
)

func main() {
	fmt.Printf("hello GOARCH=%s GOOS=%s runtime.Version()=%s\n",
		runtime.GOARCH, runtime.GOOS, runtime.Version())
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintln(os.Stderr, "ERROR: could not call debug.ReadBuildInfo(): built without Go modules?")
		return
	}
	fmt.Printf("GoVersion=%s\n", info.GoVersion)
	fmt.Printf("Module path Main.Path=%s\n", info.Main.Path)
	for _, setting := range info.Settings {
		fmt.Printf("%s=%s\n", setting.Key, setting.Value)
	}
}
