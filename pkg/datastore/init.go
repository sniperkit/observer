package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"os"
)

var db *gorm.DB

func init() {

	var err error

	// TODO: вынести в параметры
	db, err = gorm.Open("postgres",  os.Getenv(("DBCONNECTION")))
	if err != nil {

		// TODO: логирование
		fmt.Println(err.Error())
	}

	// db.LogMode(true)
	db.AutoMigrate(StackTag{}, StackQuestion{}, Settings{}, Tag{})
}