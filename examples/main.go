package main

import (
	"github.com/ginx-contribs/dbx"
	"log"
)

func main() {
	sqldb, err := dbx.Open(dbx.Options{
		Driver:   dbx.Sqlite,
		Database: "test.db",
	})

	if err != nil {
		log.Fatal(err)
	}

	sqldb.Ping()
}
