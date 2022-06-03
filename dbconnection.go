package main

import (
	// "gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() (db *gorm.DB, err error) {
	// db, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:8889)/telu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}
