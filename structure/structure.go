package structure

import (
	"fmt"

	"github.com/henriquecursino/desafioQ2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dns := "teste:teste@tcp(127.0.0.1:3306)/desafioQ2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
	}

	db.AutoMigrate(&model.Status{})
	db.AutoMigrate(&model.Categories{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Transactions{})

	fmt.Println("Connect success!")

	return db
}
