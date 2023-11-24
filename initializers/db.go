package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "ugers",
		AllowNativePasswords: true,
    }
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
        log.Fatal("Failed connect to DB")
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
	DB=db
}
