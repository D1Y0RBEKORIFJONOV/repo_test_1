package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "+_+diyor2005+_+"
	dbname   = "n9_go"
)

type Genered struct {
	ID      int
	Genered int
}

func Connect() *sql.DB {
	info := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS large_dataset (id SERIAL PRIMARY KEY,generated INT)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func InsertToTable(db *sql.DB) error {
	query := `INSERT INTO large_dataset (generated ) SELECT generate_series(1,10000000)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db := Connect()
	defer db.Close()
	err := CreateTable(db)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = InsertToTable(db)
	rows, err := db.QueryContext(ctx, "SELECT * FROM large_dataset")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var d1 []Genered
	for rows.Next() {
		var g Genered
		err = rows.Scan(&g.ID, &g.Genered)
		if err != nil {
			panic(err)
		}
		d1 = append(d1, g)
	}
	fmt.Println(d1)
	for _, g := range d1 {
		fmt.Println("ID: ", g.ID)
		fmt.Println("Genered: ", g.Genered)
	}
}
