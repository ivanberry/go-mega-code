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

	users := []model.User{
		{
			Username:     "shirting",
			PasswordHash: model.GeneratePasswordHash("abc1234"),
			Posts: []model.Post{
				{Body: "Beautiful day in Portalnd!"},
			},
		},
		{
			Username:     "谢尔婷",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Posts: []model.Post{
				{Body: "The Avengers movie was so cool!"},
				{Body: "Sun shine is beautiful"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
