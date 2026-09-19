package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("切片头本身大小: %d 字节\n\n", unsafe.Sizeof([]int{}))

	var s []int
	prevCap := 0
	fmt.Println("len    cap")
	for i := 0; i < 1200; i++ {
		s = append(s, i)
		if cap(s) != prevCap {
			fmt.Printf("%-6d %-6d\n", len(s), cap(s))
			prevCap = cap(s)
		}
	}
}
