package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Conn()  *gorm.DB{

	Db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print("koneksi databasae sukses")
	}

	return Db

	//Db.AutoMigrate(models.People{}, models.Animal{})
	//
	//Db.CreateTable(models.People{}, models.Animal{})
}

func GetDb() *gorm.DB {
	if db == nil {
		log.Print("create new conn")
		return Conn()
	}
	return db
}