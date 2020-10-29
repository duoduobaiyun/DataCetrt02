package main

import (
	"DataCertProject_Me/blockchain"
	"DataCertProject_Me/db_mysql"
	_ "DataCertProject_Me/routers"
	"fmt"
	"github.com/astaxie/beego"
)

var BUCKET_NAME ="blocks"
func main() {
	//1、实例化一个区块链实例
	bc := blockchain.NewBlockChain()
	fmt.Printf("创世区块的Hash值:%x\n", bc.LastHash)
	block, err := bc.SaveData([]byte("这里存储上链的数据信息"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("区块的高度:%d\n", block.Height)
	fmt.Printf("区块的PrevHash:%x\n", block.Hash)
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

