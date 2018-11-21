package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
)

var db *gorm.DB //database

func init() {

	e := gotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbUri := ""
	if os.Getenv("DB_PASSWORD") == "" {
		dbUri = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"))
	} else {
		dbUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"))
	}

	fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

}

func GetDB() *gorm.DB {
	return db
}
