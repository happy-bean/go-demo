package main

import "fmt"

//Go 语言条件语句
//条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

//Go 语言 if 语句
//
//if 语句由布尔表达式后紧跟一个或多个语句组成。
//语法
//Go 编程语言中 if 语句的语法如下：
//if 布尔表达式 {
///* 在布尔表达式为 true 时执行 */
//}

//Go 语言 if...else 语句
//if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
//语法
//Go 编程语言中 if...else 语句的语法如下：
//if 布尔表达式 {
///* 在布尔表达式为 true 时执行 */
//} else {
///* 在布尔表达式为 false 时执行 */
//}

//Go 语言 if 语句嵌套
//你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
//语法
//Go 编程语言中 if...else 语句的语法如下：
//if 布尔表达式 1 {
//    /* 在布尔表达式 1 为 true 时执行 */
//     if 布尔表达式 2 {
//   /* 在布尔表达式 2 为 true 时执行 */
//    }
//}

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n" )
	}
	fmt.Printf("a 的值为 : %d\n", a)

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n" );
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n" );
	}
	fmt.Printf("a 的值为 : %d\n", a);


	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
		}
	}
	fmt.Printf("a 值为 : %d\n", a );
	fmt.Printf("b 值为 : %d\n", b );
}