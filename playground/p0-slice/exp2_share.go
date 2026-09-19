package main

import "fmt"

// 共享底层数组
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b = append(b, 99)

	fmt.Println("a        =", a)
	fmt.Println("b        =", b)
	fmt.Println("len(a)   =", len(a), " cap(a) =", cap(a))
	fmt.Println("len(b)   =", len(b), " cap(b) =", cap(b))

	// 对照：用完整切片表达式限制 cap
	c := a[1:3:3]
	c = append(c, 99)
	fmt.Println("c        =", c)
	fmt.Println("a (之后) =", a)
}
