package main

import (
	"bytes"
	"fmt"
)

func main()  {
	a:=[]byte("a")
	b:=[]byte("b")
	c:=[]byte("c")

	d:=[]byte("0")
	fmt.Println(a,b,c,d)

	rs:=bytes.Join([][]byte{
		a,
		b,
		c,

	},[]byte{})
  fmt.Println("无规则拼接后的结果",rs)


	rs1:=bytes.Join([][]byte{
		a,b,c,
	},[]byte("0"))
	fmt.Println("使用0规则拼接后的结果:",rs1)
}
