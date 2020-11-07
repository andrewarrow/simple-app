package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	i := 0
	db, _ := sql.Open("mysql", "root:123456@/")
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	for {
		fmt.Println("hi-test2", i)
		i++
		time.Sleep(time.Second * 5)
	}
}
