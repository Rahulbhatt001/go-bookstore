package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm helps to talk with mysql lite and other databases

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Mysql@123@tcp(localhost:9010)/simplerest?charset= utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
