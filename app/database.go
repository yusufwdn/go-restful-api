package app

import (
	"database/sql"
	"time"
	"yusufwdn/golang-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_restful_api?tls=skip-verify&autocommit=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute) // kalau idle dalam 10 menit close aja

	return db
}
