package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "+_+diyor2005+_+"
	dbname   = "n9_go"
)

type Product struct {
	Product_name  string
	Price         float64
	category_name string
	description   string
}

func Connect() (*sql.DB, error) {
	var err error
	dataBaseINfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dataBaseINfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	product := []Product{}
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	quary := `
	SELECT p.product_name,p.price,c.category_name,c.description  FROM products p JOIN categories c ON p.category_id = c.category_id WHERE c.category_name = 'Beverages';
`
	rows, err := db.Query(quary)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var p Product
		rows.Scan(&p.Product_name, &p.Price, &p.category_name, &p.description)
		product = append(product, p)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	for _, product := range product {

		fmt.Printf("Product: %+v\t\tPrice: %+v\t\tCategory_name: %+v\t\tDescription: %+v\n", product.Product_name, product.Price, product.category_name, product.description)
	}

}
