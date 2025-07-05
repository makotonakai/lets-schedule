package database

import (
	"os"
	"fmt"
  "gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/rs/zerolog"
	gormzerolog "github.com/vitaliy-art/gorm-zerolog"
)

func Connect() *gorm.DB {

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

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

  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger,
	})

	if err != nil {
		fmt.Printf("Connection with database failed: %v", err)
	} 
	return db
}

