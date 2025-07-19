package database

import (
	"os"
	"fmt"
  "gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/rs/zerolog"
	gormzerolog "github.com/vitaliy-art/gorm-zerolog"
)

var DbUser string
var DbPassword string
var DbHost string
var DbPort string
var DbName string

func Connect() *gorm.DB {

	runLogFile, _ := os.OpenFile(
		"logs/db.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	zeroLogger := zerolog.New(multi).With().Timestamp().Logger()

	logger := gormzerolog.NewGormLogger().WithInfo(func() gormzerolog.Event {
    return &gormzerolog.GormLoggerEvent{Event: zeroLogger.Info()}
	})

  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger,
	})

	if err != nil {
		fmt.Printf("Connection with database failed: %v", err)
	} 
	return db
}

