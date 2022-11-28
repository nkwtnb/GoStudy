package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました: %v", err)
		os.Exit(1)
	}
	fmt.Println("hello from docker")
	connect()
}


func connect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(user, password,host,dbName)
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	dsn := "go_user:go_user@tcp(mysql:3306)/go_tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// fmt.Println(db, err)

	database, _ := goose.OpenDBWithDriver("mysql", dsn)
	// fmt.Println(err)
	// v, _ := goose.EnsureDBVersion(database)
	m, e := goose.CollectMigrations("db/migrations", -1, goose.MaxVersion)
	fmt.Println(e)
	fmt.Println(m)
	// 全マイグレーション数
	fmt.Println(m.Len())
	// 現在のVersion
	v, _ := goose.GetDBVersion(database)
	fmt.Println(v)
}
