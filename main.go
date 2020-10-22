package main

import (
	"DataCertProject_Me/blockchain"
	"DataCertProject_Me/db_mysql"
	_ "DataCertProject_Me/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {


	//生成第一区块
    block:=blockchain.NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
    ///fmt.Println(block)
    fmt.Printf("区块的Hash值:%x\n", block.Hash)
	fmt.Printf("区块的Hash的长度:%x\n", len(block.Hash))
	fmt.Printf("区块的Nonce值:%d\n",block.Nonce)
	return






	//1.链接数据库
	db_mysql.ConnectDB()
	//2.静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	//3.允许
	beego.Run()//启动端口监听:阻塞
}

