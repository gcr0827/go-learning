package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

func main() {
	// 指针
	fmt.Println("--------------- pointers -------------------")
	var p *int
	fmt.Printf("变量 p 自身的地址: %p\n", &p) // 打印 p 自己的门牌号
	fmt.Printf("变量 p 里面存的内容: %v\n", p) // 打印 p 里面装的值（即指向的目标）

	i := 42
	p = &i // 使用&获取变量的地址，并将地址赋值给p,即让p指向i
	fmt.Println("p:", p)

	fmt.Println(*p) // 通过指针 p 读取 i
	*p = 22         // 通过指针 p 设置 i

	fmt.Println("p:", p)
	fmt.Println("i:", i)

	// 结构体
	fmt.Println("--------------- structs -------------------")

	var u Vertex
	u = Vertex{1, 2, "test"}
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)

	fmt.Println("value", u.X)

	fmt.Println("--------------- structs-pointers -------------------")

	ps := &u
	fmt.Println("ps:", ps)
	ps.S = "UPDATE"
	fmt.Println("u:", u)
	fmt.Println("ps:", ps)

	fmt.Println("--------------- structs-字面量 -------------------")

	var (
		u1 = Vertex{1, 2, "test1"}
		u2 = Vertex{}
		u3 = &Vertex{1, 2, "test3"}
	)

	fmt.Println(u1, u2, u3)

	// 数组
	fmt.Println("--------------- array -------------------")
	var a [2]string
	a[0] = "a"
	a[1] = "b"
	fmt.Println(a)

	// 切片
	fmt.Println("--------------- slice -------------------")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[1:4]
	var s1 = primes[1:6]
	fmt.Println(s, s1)

	fmt.Println("--------------- Slices are like references to arrays -------------------")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	namea := names[0:2] // John Paul
	nameb := names[1:3] // Paul George
	fmt.Println(namea, nameb)
	nameb[0] = "xxxx"
	fmt.Println(namea, nameb)

	// 切片的长度和容量
	fmt.Println("--------------- Slices length and capacity -------------------")
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	s2 = s2[:5]
	printSlice(s2)

	s2 = s2[:4]
	printSlice(s2)

	s2 = s2[2:]
	printSlice(s2)

	fmt.Println("--------------- Slice empty nil-------------------")

	var es []int
	fmt.Println(es, len(es), cap(es))
	fmt.Println(&es)
	fmt.Printf("es: %#v, 指针: %p\n", es, es) // s1: []int(nil), 指针: 0x0

	fmt.Println("--------------- create slice with make -------------------")
	ma := make([]int, 5)
	printSlice(ma)

	mb := make([]int, 0, 5)
	printSlice(mb)
	fmt.Printf("es: %#v, 指针: %p\n", mb, mb)

	mc := mb[:2]
	printSlice(mc)

	md := mc[2:5]
	printSlice(md)

	fmt.Println("--------------- range -------------------")

	var powa = []int{2, 3, 5, 7, 11, 13}
	for q, w := range powa {
		fmt.Printf("%d = %d\n", q, w)
	}

	fmt.Println("--------------- map -------------------")

	mm := make(map[string]Vertex)
	fmt.Println(mm)
	fmt.Printf("es: %#v, 指针: %p\n", mm, mm)

	mm["test1"] = Vertex{1, 2, "test2"}
	fmt.Println(mm)

	fmt.Println("--------------- map字面量 -------------------")

	ml := map[string]Vertex{
		"test1": Vertex{1, 2, "test2"},
		"test2": Vertex{3, 4, "test3"},
	}
	fmt.Println(ml)

	var mm2 map[string]Vertex
	fmt.Printf("mm2: %#v, 指针: %p\n", mm2, mm2)

	fmt.Println("--------------- map的操作 -------------------")

	// mm3 和 mm4功能上是等价的
	mm3 := map[string]int{} // 字面量的方式
	fmt.Printf("mm2: %#v, 指针: %p\n", mm3, mm3)

	mm3["test"] = 1

	fmt.Println(mm3)
	mm4 := make(map[string]int)
	fmt.Printf("mm2: %#v, 指针: %p\n", mm4, mm4)

	fmt.Println("--------------- function -------------------")

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println("--------------- function 作为其他函数的入参-------------------")

	fmt.Println(compute(hypot))

	fmt.Println("--------------- function闭包-------------------")
	pos, neg := adder(), adder()
	for i := 0; i < 2; i++ {
		fmt.Println(pos(i), neg(i+1))
	}

	fmt.Println("--------------- Methods-------------------")
	v1 := Vertex1{3, 4}
	fmt.Println(v1.Abs())

	fmt.Println("--------------- Methods 指针接收者-------------------")

	fr := MyFloat(3) // 显示转换为MyFloat
	fr.Double()
	fmt.Println(fr.Abs())

	fmt.Println("--------------- Interface-------------------")
	var anyValue interface{}
	anyValue = 42
	fmt.Println(anyValue)

	fmt.Println("--------------- Empty Interface-------------------")

	var ii interface{}
	describe(ii)

	ii = 42
	describe(ii)

	ii = "test"
	describe(ii)

	// 类型断言
	fmt.Println("--------------- Type assertions -------------------")

	var ia interface{} = "hello"

	is := ia.(string)
	fmt.Println(is)

	is, ok := ia.(string)
	fmt.Println(is, ok)

	fi, ok := ia.(float64)
	fmt.Println(fi, ok)

	// 类型switch
	do(21)

	fmt.Println("--------------- Stringer Interface -------------------")
	si := Person{"Arthur Dent", 42}
	sz := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(si, sz)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println("--------------- Errors -------------------")
	/*if err := run(); err != nil {
		fmt.Println(err)
	}*/

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyError struct {
	When time.Time
	What string
}

func run() error {
	return &MyError{time.Now(), "test"}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

type Vertex struct {
	X, Y int
	S    string
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Vertex1 struct {
	X, Y float64
}

func (v Vertex1) Abs() float64 {
	return math.Abs(v.X - v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f))
}

func (f *MyFloat) Double() {
	*f = *f * *f
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)

	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
