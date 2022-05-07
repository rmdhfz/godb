package godb

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/dbgo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
