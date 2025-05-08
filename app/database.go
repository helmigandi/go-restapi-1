package app

import (
	"database/sql"
	"github.com/helmigandi/go-restapi-1/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/mydb")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
