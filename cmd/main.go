package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rkunihiro/gormgql/repository"
)

func printJSON(v any) {
	buf, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		log.Println(string(buf))
	}
}

func main() {
	conn, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/dbname?parseTime=true")
	if err != nil {
		panic(fmt.Errorf("%v", err))
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{Conn: conn}),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	user, err := userRepo.FindByID(1)
	if err != nil {
		panic(err)
	}
	printJSON(user)

	users, err := userRepo.Find()
	if err != nil {
		panic(err)
	}
	printJSON(users)

	post, err := postRepo.FindByID(1)
	if err != nil {
		panic(err)
	}
	printJSON(post)

	posts, err := postRepo.Find()
	if err != nil {
		panic(err)
	}
	printJSON(posts)
}
