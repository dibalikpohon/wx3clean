package db

import (
	"database/sql"
	"fmt"
	"time"
	"wx3clean/internal/config"
)

var db_ *sql.DB

func MysqlGetInstance() *sql.DB {
	if db_ == nil {
		panic("mysql is not open!")
	}
	return db_
}

func MysqlEndInstance() {
	if db_ != nil {
		db_.Close()
		db_ = nil
	}
}

func MysqlOpen() {

	appConfig := config.AppGetConfig()

	username := appConfig.MysqlDB_USER
	password := appConfig.MysqlDB_PASS
	host := appConfig.MysqlDB_HOST
	port := appConfig.MysqlDB_PORT
	db := appConfig.MysqlDB_DB

	connetionstring := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	db_, err := sql.Open("mysql", fmt.Sprintf(connetionstring, username, password, host, port, db))
	if err != nil {
		panic(err)
	}

	db_.SetConnMaxIdleTime(10 * time.Minute)
	db_.SetConnMaxLifetime(12 * time.Hour)
	db_.SetMaxIdleConns(10)
	db_.SetMaxOpenConns(100)

	err = db_.Ping()
	if err != nil {
		panic(err)
	}
}
