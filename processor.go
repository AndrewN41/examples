// processor prints info about the architecture it is running on
// 	Operating System, Architecture, Integer Size, Pointer Size
//
package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	const PtrSize = 32 << uintptr(^uintptr(0)>>63)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("Integer:", strconv.IntSize)
	fmt.Println("Pointer:", PtrSize)
}
