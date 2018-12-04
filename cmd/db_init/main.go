package main

import (
	"fmt"
	"log"

	"github.com/ivanberry/go-mega-code/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init...")
	db := model.ConnectToDB()

	fmt.Printf("the db url is %p\n", db)
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})
}
