package model

import (
	"fmt"
	"log"

	"github.com/ivanberry/go-mega-code/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

//ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectStr := config.GetMysqlConnectingString()
	fmt.Printf("database url %s\n", connectStr)
	log.Println("Connet to db...")
	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		fmt.Printf("ERROR %v\n", err)
		panic("failed to connect database")
	}
	db.SingularTable(true)
	return db
}
