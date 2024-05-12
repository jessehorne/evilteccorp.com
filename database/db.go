package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"
	"fmt"
	"os"
)

var GDB *gorm.DB

func GetDSN() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		user, pass, host, port, name)
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", GetDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitGDB() error {
	db, err := gorm.Open(mysql.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	GDB = db

	return nil
}
