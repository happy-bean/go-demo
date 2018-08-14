package main

import "fmt"

//Go 语言运算符
//运算符用于在程序运行时执行数学或逻辑运算。
//Go 语言内置的运算符有：
//算术运算符
//关系运算符
//逻辑运算符
//位运算符
//赋值运算符
//其他运算符
//接下来让我们来详细看看各个运算符的介绍。
//参考：http://www.runoob.com/go/go-operators.html

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c )
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c )
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c )
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c )
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a )
	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )
}

//关系运算符
//下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20
//
//逻辑运算符
//下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。
//
//位运算符
//位运算符对整数在内存中的二进制位进行操作。
//下表列出了位运算符 &, |, 和 ^ 的计算：
//
//赋值运算符
//下表列出了所有Go语言的赋值运算符。
//
//运算符优先级
//有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

//参考：http://www.runoob.com/go/go-operators.html