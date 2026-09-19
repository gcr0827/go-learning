package main

import "fmt"

// 实验3：传参的拷贝语义
func modify(s []int) {
	s[0] = 999
	s = append(s, 1000)
	fmt.Println("  函数内 s =", s, " len =", len(s), " cap =", cap(s)) // [999 2 3 1000]  len = 4  cap = 10
}

func modifyPtr(s *[]int) {
	(*s)[0] = 888
	*s = append(*s, 2000)
}

func main() {
	s := make([]int, 3, 10)
	s[0], s[1], s[2] = 1, 2, 3

	fmt.Println("初始     s =", s, " len =", len(s), " cap =", cap(s)) // [1,2,3] 3, 10
	modify(s)
	fmt.Println("值传递后 s =", s, " len =", len(s), " cap =", cap(s)) // [999,2,3] 3, 10

	modifyPtr(&s)
	fmt.Println("指针传递后 s =", s, " len =", len(s), " cap =", cap(s)) // [888, 2,3,4000,5,10]
}
