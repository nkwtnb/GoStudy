package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

func main() {
	fmt.Println("hello from docker")
	connect()
}


func connect() {
	// 各環境変数は /.env に記載し、docker-compose.ymlで読み込む想定
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(user, password,host,dbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, _ := goose.OpenDBWithDriver("mysql", dsn)
	m, e := goose.CollectMigrations("db/migrations", -1, goose.MaxVersion)
	fmt.Println(e)
	fmt.Println(m)
	// 全マイグレーション数
	fmt.Println(m.Len())
	// 現在のVersion
	v, _ := goose.GetDBVersion(db)
	fmt.Println(v)
}
