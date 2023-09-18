package configs

import (
	"fmt"
	itemDB "gorest-10/models/item/database"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("db_user"), os.Getenv("db_password"), os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_name"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal init database")
	}
	Migration()
}

// migration
func Migration() {
	DB.AutoMigrate(itemDB.Item{})
}
