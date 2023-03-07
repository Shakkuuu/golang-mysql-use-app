package db

import (
	"fmt"
	"time"

	"db-app/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init(un string, up string, dbn string) {
	fmt.Println("db接続開始")

	DBMS := "mysql"
	USER := un
	PASS := up
	PROTOCOL := "tcp(db:3309)"
	DBNAME := dbn

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Printf("db Init error: %v\n", err)
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		fmt.Printf("db Close error: %v\n", err)
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{}, &entity.Coin{}, entity.Buki{})
}
