package main

import (
	"fmt"
)

func main() {
	fmt.Println(PlusOne(1))
}

//go:wasmexport plus_one
func PlusOne(num int32) int32 {
	return num + 1
}
