package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
)

var gormdb *gorm.DB //database

func init() {

	e := gotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbURI := ""
	if os.Getenv("DB_PASSWORD") == "" {
		dbURI = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"))
	} else {
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"))
	}

	fmt.Println(dbURI)

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	gormdb = conn

}

func GetGormDB() *gorm.DB {
	return gormdb
}
