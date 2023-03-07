package main

import (
	"fmt"
	"net/http"
	"os"

	"db-app/db"
	"db-app/entity"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("serverスタート")

	un, up, dbn := loadEnv()

	db.Init(un, up, dbn)
	// server.Init()
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello from h1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		a, b := GetAllUsers()
		if b != nil {
			fmt.Println(b)
			return
		}
		fmt.Fprintln(w, a)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/h2", h2)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}

	db.Close()
}

func loadEnv() (string, string, string) {
	fmt.Println("envファイル読み込み開始")
	err := godotenv.Load("db-api.env")

	if err != nil {
		fmt.Printf("envファイルの読み込みエラー: %v\n", err)
	}

	username := os.Getenv("USERNAME")
	userpass := os.Getenv("USERPASS")
	database := os.Getenv("DATABASE")

	// fmt.Println(username)
	// fmt.Println(userpass)
	// fmt.Println(database)

	fmt.Println("envファイル読み込み完了")

	return username, userpass, database
}

// GetAllUsers is get all Users
func GetAllUsers() ([]entity.User, error) {
	que := db.GetDB()
	var u []entity.User

	if err := que.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
