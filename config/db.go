package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	tempDB, err := sql.Open("mysql", "root:root@/gocrud")

	if err != nil {
		panic(err)
	}

	if err := tempDB.Ping(); err != nil {
		panic(err)
	}

	db = tempDB
	fmt.Println("Database connected!")
}

func Initialize() {
	if err := CreateUsersTable(); err != nil {
		// Handle error
		panic(err)
	}
}

func CreateUsersTable() error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS user (
			id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            is_verified BOOLEAN NOT NULL,
            role INT NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
