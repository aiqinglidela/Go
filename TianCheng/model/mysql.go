package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL()(err error){
	//配置MySQL连接参数
	username := "root"  //账号
	password := "aimeng1321" //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306 //数据库端口
	Dbname := "tiancheng" //数据库名
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	DB, err = gorm.Open("mysql", dsn)
	//设置数据库连接池参数
	DB.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	DB.DB().SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return DB
}

func Close(){
	DB.Close()
}