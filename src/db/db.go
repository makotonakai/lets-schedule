package db

import (
	"os"
	"fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func Connect() *gorm.DB {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname:= os.Getenv("DB_NAME")

  dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection failed: %s", err)
	}

	return db
}