package dbconnect

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB{
db, err := gorm.Open("mysql", "DBへの接続")
    if err != nil {
        fmt.Print("データベース接続に失敗しました。")
    }
    db.LogMode(true)
    return db
}