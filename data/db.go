package data

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xyingsoft/golang-vue/config"
)

// 	"database/sql"
// 	"os"
// 	_ "github.com/go-sql-driver/mysql"

// var mainDB *sql.DB
// func DB() *sql.DB { return mainDB }

var mainDB *gorm.DB
var RecordNotFound = errors.New("Record not found.")
var NullRecord = errors.New("Null record.")
var initer sync.Once

func DB() *gorm.DB { return mainDB }

func initializeDatabase() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	var url string

	url = config.GetVar(config.VAR_DATABASE_URL)
	db, err = gorm.Open("mysql", url)

	fmt.Println("Database connecting...")
	fmt.Println(url)

	if err != nil {
		panic(err)
	}

	return db, err
}

func initDatabases() {
	db, err := initializeDatabase()
	if err != nil {
		panic(err)
	}
	mainDB = db
}

func Init() {
	initer.Do(func() {
		initDatabases()
	})
}
