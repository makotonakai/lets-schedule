package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}
	
	gormDB, err := gorm.Open("mysql", db) //通常ならここにdbのurlが入るが、テストの場合はmockで作ったdbを入れる。
	if err != nil {
		return gormDB, mock, err
	}
	return gormDB, mock, err
}

