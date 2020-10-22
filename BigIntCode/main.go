package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 1
	b := 3
	fmt.Println(a+b, a == b)
	fmt.Println()
	var a1 int
	a1 = 3
	var b1 int
	b1 = 5
	fmt.Println(a1+b1, a != b)

	//内置函数,新建一个big.Int类型,并初始化,返回一个指针
	a2 := big.NewInt(1)
	fmt.Println(a2)
	b2 := big.NewInt(3)
	//a2 = a2.Add(a2, b2)
	//比较cmp
	//一共有三个返回值-1,0,1

	rs:=a2.Cmp(b2)
	fmt.Println("a2和b2的大小关系",rs)//a2<b2 = -1 , a2==b2 =0 ,a2>b2=1
	fmt.Println("两个大整数相加", a2)

	a3:=big.NewInt(1)
	a3.Lsh(a3,1)
	fmt.Println("a3左移一位后",a3)

	a4:=big.NewInt(5)
	a4.Rsh(a4,5)
	fmt.Println("a4右移一位后",a4)
}
