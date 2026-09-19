package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java = true, false, "no"

func main() {

	fmt.Println("Hello World GCR")
	fmt.Println("-----------------------")

	fmt.Println("Welcome to the playground")
	fmt.Println("The time is", time.Now())
	fmt.Println("-----------------------")

	// package
	fmt.Println("------------- Package -------------")

	// rand => 随机数包
	// Intn 返回非负随机整数，范围 0 到 MaxInt =》 前闭后开
	fmt.Println("My favorite number is", rand.Intn(10))

	// math => 数学函数包，sqrt 开平方
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))

	fmt.Println(math.Pi)

	// 函数
	fmt.Println("------------- Function -------------")
	fmt.Println("调用相加的函数", funcAdd(1, 2))
	fmt.Printf("调用相加的函数,结果为:%d \n", funcAdd(1, 2))

	// 函数返回多结果
	x, y := swap("hello", "world")
	fmt.Println("函数返回多个值的结果", x, y)

	a, b := split(10)
	fmt.Println("函数返回多个结果值有命名", a, b)

	// 函数
	fmt.Println("------------- Variable -------------")
	var c, python, java = true, false, "yes"
	fmt.Println("变量的输出:", c, python, java)

	// 类型转换
	fmt.Println("------------- Convert -------------")
	var i int = 42
	var f float64 = float64(i)
	fmt.Println(i, f)

	v := int(42.0)
	fmt.Printf("v is of type %T\n", v)

	fmt.Println("------------- 循环 -------------")
	sum := 0
	for q := 0; q < 10; q++ {
		sum += q
	}
	fmt.Println("循环求和", sum)

	// while 语句 for的写法
	sumWhile := 1
	for sumWhile < 10 {
		// 1 +1 2
		// 2+2 =4
		// 4+4=8
		// 8+8 =16
		sumWhile += sumWhile
	}
	fmt.Println("while-for的写法", sumWhile)

	// if 循环
	fmt.Println("------------- if循环 -------------")
	fmt.Println("if 循环", sqrt(9))

	fmt.Println("------------- if循环 短语句-------------")
	fmt.Println("if 短语句", pow(2, 3, 7))

	fmt.Printf("%g\n", 123.456) // 123.45
	fmt.Printf("%f\n", 123.456) // 123.456000

	fmt.Println("------------- switch-------------")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When Saturday")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	fmt.Println("Switch no condition")
	t := time.Now()
	fmt.Println(t.Hour(), t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 延迟执行
	fmt.Println("------------- defer-------------")
	defer fmt.Println("defer test")

	// 延迟堆叠，使用栈，后进先出
	/*for k := 0; k < 2; k++ {
		defer fmt.Println(k)
	}*/

	for k := 0; k < 2; k++ {
		func(k int) { // 1. 定义了一个匿名函数
			defer fmt.Println(k) // 这个 defer 属于这个匿名函数，不属于 outer！
		}(k) // 2. 在这里立即调用这个匿名函数
		fmt.Println("test")
	}

}

func funcAdd(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
