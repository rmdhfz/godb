package godb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

// func TestQuerySQL(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	parent := context.Background()
// 	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
// 	defer cancel()

// 	script := "SELECT * FROM customer"
// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("ID", id)
// 		fmt.Println("Name", name)
// 	}
// 	fmt.Println("success")
// }

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
	var balance sql.NullInt32
	var rating sql.Float64
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
		if balance.Valid {
			fmt.Println("Balance", balance.Int32)
		}
		if rating.Valid {
			fmt.Println("Rating", rating.Float64)
		}
		fmt.Println("Married", married)
		fmt.Println("Created At", createdAt)
	}
	fmt.Println("success")
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()
	username := "admin'; #"
	password := "salah"
	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()
	username := "admin"
	password := "admin"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExecLastInsertID(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	username := "test"
	password := "test"

	script := "INSERT INTO user(username, password) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert ID", insertID)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	script := "INSERT INTO user (username, password) VALUES(?,?)"
	statement, err := db.PrepareContext(parent, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		username := "test" + strconv.Itoa(i)
		password := "test" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, username, password)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO user (username, password) VALUES(?,?)"
	// transaksi
	for i := 0; i < 10; i++ {
		username := "test" + strconv.Itoa(i)
		password := "test" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, username, password)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
