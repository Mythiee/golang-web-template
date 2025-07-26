package db

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var DB *bun.DB

func Init() {
	dbDriverName := os.Getenv("DB_DRIVER_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open(dbDriverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	// bun.DB
	DB = bun.NewDB(db, sqlitedialect.New())

	if err := DB.Ping(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
