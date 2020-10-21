package main

import (
	"DataCertProject_Me/db_mysql"
	_ "DataCertProject_Me/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1.链接数据库
	db_mysql.ConnectDB()
	//2.静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	//3.允许
	beego.Run()//启动端口监听:阻塞
}

