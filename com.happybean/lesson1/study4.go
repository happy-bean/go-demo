package main

//Go 语言变量
//变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。变量可以通过变量名访问。
//Go 语言变量名由字母、数字、下划线组成，其中首个字母不能为数字。
//声明变量的一般形式是使用 var 关键字：
//var identifier type
//	变量声明
//第一种，指定变量类型，声明后若不赋值，使用默认值。
//var v_name v_type
//v_name = value
//第二种，根据值自行判定变量类型。
//var v_name = value
//第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
//v_name := value

// 例如
//var a int = 10
//var b = 10
//c :=10

var a = "菜鸟教程"
var b string = "runoob.com"
var c bool

func main(){
	d := 0
	println(a, b, c, d)
}

//多变量声明
//类型相同多个变量, 非全局变量
//var vname1, vname2, vname3 type
//	vname1, vname2, vname3 = v1, v2, v3
//
//var vname1, vname2, vname3 = v1, v2, v3 //和python很像,不需要显示声明类型，自动推断
//
//vname1, vname2, vname3 := v1, v2, v3 //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误
//
//
//// 这种因式分解关键字的写法一般用于声明全局变量
//var (
//	vname1 v_type1
//	vname2 v_type2
//)

//var x, y int
//var (  // 这种因式分解关键字的写法一般用于声明全局变量
//	a int
//	b bool
//)
//
//var c, d int = 1, 2
//var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"