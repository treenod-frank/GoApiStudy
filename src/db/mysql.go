package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var orm *gorm.DB

func OpenDatabase() *gorm.DB {
	dsn := "treenod:xmflshem@tcp(ec2-13-125-229-201.ap-northeast-2.compute.amazonaws.com:3306)/study"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func SetupDatabase() {
	db := OpenDatabase()
	orm = db
}

func GetORM() *gorm.DB {
	return orm
}
