package godb

import (
	"context"
	"database/sql"
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

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	script := "SELECT id, name, email, balance, rating, brith_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id, name string
	var email sql.NullString
	var balance int32
	var rating float64
	var brithDate, createdAt time.Time
	var married bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &balance, &rating, &brithDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID", id)
		fmt.Println("Name", name)
		if email.Valid {
			fmt.Println("Email", email.String)
		}
		fmt.Println("Balance", balance)
		fmt.Println("Rating", rating)
		fmt.Println("Married", married)
		fmt.Println("Created At", createdAt)
	}
	fmt.Println("success")
}
