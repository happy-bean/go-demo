package main

import "fmt"

//Go 语言变量作用域
//作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。
//Go 语言中变量可以在三个地方声明：
//函数内定义的变量称为局部变量
//函数外定义的变量称为全局变量
//函数定义中的变量称为形式参数

//局部变量
//在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
//以下实例中 main() 函数使用了局部变量 a, b, c：

//全局变量
//在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
//全局变量可以在任何函数中使用，以下实例演示了如何使用全局变量：

/* 声明全局变量 */
var g int

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a);
	fmt.Printf("sum() 函数中 b = %d\n", b);
	return a + b;
}

func main() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("结果： a = %d, b = %d and c = %d and g = %d\n", a, b, c, g)

	d := sum(a, b);
	fmt.Printf("main()函数中 d = %d\n", d);
}
