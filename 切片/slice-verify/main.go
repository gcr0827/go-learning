// 切片行为全量验证程序（Go 1.26.5 windows/amd64 本机实测）
// 运行：cd D:\Program\go-learning\slice-verify && go run .
package main

import (
	"fmt"
	"unsafe"
)

func headerSize() {
	fmt.Println("=== 1. slice header 大小 ===")
	fmt.Printf("  unsafe.Sizeof([]int{})  = %d 字节\n", unsafe.Sizeof([]int{}))
	fmt.Printf("  unsafe.Sizeof([]byte{}) = %d 字节\n", unsafe.Sizeof([]byte{}))
	fmt.Println("  -> ptr(8) + len(8) + cap(8) = 24（64 位平台）")
}

func initForms() {
	fmt.Println("\n=== 2. 不同初始化方式的 len/cap ===")
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := make([]int, 3)
	s4 := make([]int, 3, 10)
	s5 := make([]int, 0, 10)
	fmt.Printf("  var s []int          len=%-3d cap=%-3d  nil=%v\n", len(s1), cap(s1), s1 == nil)
	fmt.Printf("  []int{1,2,3}         len=%-3d cap=%-3d\n", len(s2), cap(s2))
	fmt.Printf("  make([]int, 3)       len=%-3d cap=%-3d\n", len(s3), cap(s3))
	fmt.Printf("  make([]int, 3, 10)   len=%-3d cap=%-3d\n", len(s4), cap(s4))
	fmt.Printf("  make([]int, 0, 10)   len=%-3d cap=%-3d  <- 预分配常用\n", len(s5), cap(s5))
}

func growthInt() {
	fmt.Println("\n=== 3. []int 逐个 append 的 cap 序列（实测）===")
	var s []int
	prev := 0
	for i := 0; i < 3500; i++ {
		s = append(s, i)
		if cap(s) != prev {
			fmt.Printf("  len=%-5d cap=%-6d\n", len(s), cap(s))
			prev = cap(s)
		}
	}
}

func growthByte() {
	fmt.Println("\n=== 4. []byte 逐个 append 的 cap 序列（实测）===")
	var b []byte
	prev := 0
	for i := 0; i < 3500; i++ {
		b = append(b, 'x')
		if cap(b) != prev {
			fmt.Printf("  len=%-5d cap=%-6d\n", len(b), cap(b))
			prev = cap(b)
		}
	}
}

func formulaCheck() {
	fmt.Println("\n=== 5. nextslicecap 公式验证（runtime/slice.go:326）===")
	fmt.Println("  公式：newcap += (newcap + 3*256) >> 2")
	fmt.Println("  initCap  appendN  新len  新cap   2*cap   newLen>2cap?")
	cases := []struct{ initCap, addN int }{
		{8, 1000}, {8, 16}, {8, 17}, {8, 100}, {256, 100}, {1024, 100},
	}
	for _, c := range cases {
		x := make([]byte, c.initCap, c.initCap)
		y := append(x, make([]byte, c.addN)...)
		fmt.Printf("  %-8d %-8d %-6d %-6d %-6d %v\n",
			c.initCap, c.addN, len(y), cap(y), c.initCap*2, len(y) > c.initCap*2)
	}
}

func writePosition() {
	fmt.Println("\n=== 6. append 写入位置 = 起点 + len（不是找空位）===")
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Printf("  a = %v\n", a)
	fmt.Printf("  b = a[1:3]  len=%d cap=%d  &b[0]==&a[1]? %v\n", len(b), cap(b), &b[0] == &a[1])
	fmt.Printf("  写入位置 = 起点(idx1) + len(2) = idx3；idx3 原值 = %d（非空位）\n", a[3])
	b = append(b, 99)
	fmt.Printf("  append(b,99) 后 a = %v  <- idx3 的 4 被覆盖\n", a)
	fmt.Printf("                  b = %v\n", b)
}

func noRealloc() {
	fmt.Println("\n=== 7. cap 足够时 append 不换底层数组 ===")
	s := make([]int, 3, 10)
	before := &s[0]
	s = append(s, 1000)
	after := &s[0]
	fmt.Printf("  make([]int,3,10) + append -> len=%d cap=%d\n", len(s), cap(s))
	fmt.Printf("  &s[0] 前=%p  后=%p  相同? %v\n", before, after, before == after)
}

var outsideS []int

func inspect(s []int) {
	fmt.Printf("    [函数内] &s=%p  &s[0]=%p  len=%d cap=%d\n", &s, &s[0], len(s), cap(s))
	s[0] = 999
	s = append(s, 1000)
	fmt.Printf("    [函数内] append 后 &s=%p  len=%d cap=%d\n", &s, len(s), cap(s))
}

func valueSemantics() {
	fmt.Println("\n=== 8. 值传递证明：&s 不同、&s[0] 相同 ===")
	s := make([]int, 3, 10)
	s[0], s[1], s[2] = 1, 2, 3
	fmt.Printf("    [函数外] &s=%p  &s[0]=%p  len=%d cap=%d\n", &s, &s[0], len(s), cap(s))
	inspect(s)
	fmt.Printf("    [函数外] &s=%p  &s[0]=%p  len=%d cap=%d\n", &s, &s[0], len(s), cap(s))
	fmt.Printf("    调用后 s = %v（len 仍 3，s[0]=%d 说明元素改动生效）\n", s, s[0])
	fmt.Println("    -> &s 不同 = header 被复制（值传递）")
	fmt.Println("    -> &s[0] 相同 = 底层数组共享")
}

func sharedArrayProof() {
	fmt.Println("\n=== 9. 共享底层数组：数据写进去了但够不着 ===")
	s := make([]int, 3, 10)
	s[0], s[1], s[2] = 1, 2, 3
	inspect(s)
	fmt.Printf("    调用后 s = %v\n", s)
	fmt.Printf("    但 s[:cap(s)] = %v  <- 1000 真在底层数组里\n", s[:cap(s)])
	s = append(s, 999)
	fmt.Printf("    append(s,999) 后 s = %v  <- 999 覆盖了那个 1000\n", s)
}

func fullSliceExpr() {
	fmt.Println("\n=== 10. 完整切片表达式 a[low:high:max] ===")
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("  a[1:3]   len=%d cap=%d  (cap 到数组末尾)\n", len(a[1:3]), cap(a[1:3]))
	c := a[1:3:3]
	fmt.Printf("  a[1:3:3] len=%d cap=%d  (cap 被限制)\n", len(c), cap(c))
	c = append(c, 99)
	fmt.Printf("  append(c,99) 后 a = %v  <- 未受影响（cap 不够，分配新数组）\n", a)
	fmt.Printf("                  c = %v\n", c)
}

func main() {
	headerSize()
	initForms()
	growthInt()
	growthByte()
	formulaCheck()
	writePosition()
	noRealloc()
	valueSemantics()
	sharedArrayProof()
	fullSliceExpr()
}
