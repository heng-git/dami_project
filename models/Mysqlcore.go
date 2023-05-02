package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	//password := config.Section("mysql").Key("password").String()
	data := config.Section("mysql").Key("database").String()
	//dsn := "root:@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, ip, port, data)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
}
