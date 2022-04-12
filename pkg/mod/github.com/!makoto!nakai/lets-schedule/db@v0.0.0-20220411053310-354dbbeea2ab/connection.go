package dbconnect

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB{
db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        fmt.Print("データベース接続に失敗しました。")
    }
    db.LogMode(true)
    return db
}