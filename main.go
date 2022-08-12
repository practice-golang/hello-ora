package main // import "hello-ora"

import (
	"context"
	"database/sql/driver"
	"log"
	"os"

	go_ora "github.com/sijms/go-ora/v2"
)

func die(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		os.Exit(1)
	}
}

func main() {
	addr := "oracle://system:oracle@localhost:1521/XE"

	con, err := go_ora.NewConnection(addr)
	die("Create connection:", err)

	err = con.Open()
	die("Open connection:", err)

	defer con.Close()

	err = con.Ping(context.Background())
	die("Ping connection:", err)

	log.Println("Connected")
	stmt := go_ora.NewStmt("SELECT * FROM v$version", con)
	defer stmt.Close()

	rows, err := stmt.Query(nil)
	die("Can't create query:", err)
	defer rows.Close()

	values := make([]driver.Value, 1)

	for {
		err = rows.Next(values)
		if err != nil {
			die("Can't fetch row:", err)
		}

		log.Println(values[0])
	}
}
