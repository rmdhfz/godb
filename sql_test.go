package godb

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	script := "INSERT INTO customer(name) VALUES ('Wanda Maulidina')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("success")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID", id)
		fmt.Println("Name", name)
	}
	fmt.Println("success")
}
