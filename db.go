package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn5 := "root:rootpwd@tcp(127.0.0.1:5306)/mydb"
	testValidTime(dsn5, "MySQL 5")
	testYear2Time(dsn5, "MySQL 5")
	testYear1Time(dsn5, "MySQL 5")
	testUninitialisedTime(dsn5, "MySQL 5")
	testNullTime(dsn5, "MySQL 5")

	log.Println()

	dsn8 := "root:root@tcp(127.0.0.1:8306)/mydb"
	testValidTime(dsn8, "MySQL 8")
	testYear2Time(dsn8, "MySQL 8")
	testYear1Time(dsn8, "MySQL 8")
	testUninitialisedTime(dsn8, "MySQL 8")
	testNullTime(dsn8, "MySQL 8")
}

func testNullTime(dsn string, label string) {
	testWithTime(dsn, label, "null time", nil)
}

func testUninitialisedTime(dsn string, label string) {
	testWithTime(dsn, label, "uninitialised date", time.Time{})
}

func testYear2Time(dsn string, label string) {
	testWithTime(dsn, label, "year 2 AD", time.Date(2, 1, 1, 0, 0, 0, 0, &time.Location{}))
}

func testYear1Time(dsn string, label string) {
	testWithTime(dsn, label, "year 1 AD", time.Date(1, 1, 1, 0, 0, 0, 0, &time.Location{}))
}

func testValidTime(dsn string, label string) {
	testWithTime(dsn, label, "valid time", time.Date(2021, 1, 1, 0, 0, 0, 0, &time.Location{}))
}

func testWithTime(dsn string, label string, desc string, startDate interface{}) {
	ctx := context.Background()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	query := `SELECT * FROM person WHERE created_at >= ?`

	rows, err := db.QueryContext(ctx, query, startDate)
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		var id int64
		var name string
		var created []uint8 // ignore `?parseTime=true` support for simplicity

		err = rows.Scan(&id, &name, &created)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s: %s: result is {%d, %s, %s}", label, desc, id, name, created)
	} else {
		log.Printf("%s: %s: record not found", label, desc)
	}
}
