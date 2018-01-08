package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

var db *gorm.DB

func init() {

	var err error

	// TODO: вынести в параметры
	db, err = gorm.Open("postgres", "host=192.168.1.71 user=root dbname=rss sslmode=disable password=root")
	if err != nil {

		// TODO: логирование
		fmt.Println(err.Error())
	}

	//db.LogMode(true)
	db.AutoMigrate(StackTag{}, StackQuestion{}, Settings{}, Tag{})
}