package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"

)

var Db *sql.DB
func ConnectDB() {
	//1.读取conf配置信息
	config := beego.AppConfig
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//fmt.Println(dbDriver,dbUser,dbPassword,dbName)

	//2.组织链接数据库的字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	//3.链接数据库
	db, err := sql.Open(dbDriver, connUrl)
	//4.获取数据库链接对象,处理链接结果
	if err != nil {
		//早发现,早解决
		fmt.Println(err.Error())
		panic("数据库连接错误,请检查配置")
	}
	//为全局的数据库操作对象赋值
	Db=db
}
