package app

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	// config := mysql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   "localhost:3306",
	// 	DBName: "myPorto",
	// }
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/myPorto")

	if err != nil {
		log.Fatal(err)
	}

	if ping := conn.Ping(); ping != nil {
		log.Fatal(ping.Error())
	}
	conn.SetConnMaxIdleTime(60 * time.Second)
	conn.SetConnMaxLifetime(60 * time.Second)
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(10)

	return conn
}
