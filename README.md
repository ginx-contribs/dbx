# dbx

dbx is a database/sql helper, support follow driver:

* mysql
* postgres
* sqlite
* sqlserver

## install
```bash
go get github.com/ginx-contribs/dbx@latest
```


## usage

here is a sqlite example as follows
```go
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
```
