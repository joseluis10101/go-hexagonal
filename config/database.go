package config

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQLConnection() *sqlx.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		"root",
		"mysql",
		"127.0.0.1",
		3306,
		"inspector",
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Panic("error connecting to database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(15 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Panic("error pinging database: %w", err)
	}

	return db
}
