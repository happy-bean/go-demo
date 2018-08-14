package main

import "fmt"

//iota
//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
//在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
//iota 可以被用作枚举值：
//const (
//	a = iota
//	b = iota
//	c = iota
//)

//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
//const (
//	a = iota
//	b
//	c
//)

const (
	ii=1<<iota
	j=3<<iota
	k
	l
)

func main() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)

	fmt.Println("ii=",ii)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}

//iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
//简单表述:
//i=1：左移 0 位,不变仍为 1;
//j=3：左移 1 位,变为二进制 110, 即 6;
//k=3：左移 2 位,变为二进制 1100, 即 12;
//l=3：左移 3 位,变为二进制 11000,即 24。