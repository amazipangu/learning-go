package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capure connection properties
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("MARIADB_ROOT_USER")
	cfg.Passwd = os.Getenv("MARIADB_ROOT_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = "localhost:3306"
	cfg.DBName = "recordings"

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
