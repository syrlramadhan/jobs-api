package config

import (
	"database/sql"
	"geemod/jobs-api/helper"
	"time"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/corporate?parseTime=true")
	helper.SendPanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
