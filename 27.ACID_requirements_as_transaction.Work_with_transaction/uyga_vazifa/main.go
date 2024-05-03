package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	dbname   = "n9_go"
	user     = "postgres"
	port     = 5432
	password = "+_+diyor2005+_+"
)

type DB struct {
	DB *sql.DB
	Tx *sql.Tx
}

type Categories struct {
	Category_id   uint
	Category_name string
	Description   string
}

type Product struct {
	Product_id   uint
	Product_name string
	Category_id  uint
	Unit         string
	Price        float64
	Category     Categories
}

func ConnectDb() (*DB, error) {
	strInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err := sql.Open("postgres", strInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func (db *DB) InsertIntoProduct(product_name, unit string, category_id uint, price float64) (product Product, err error) {

	query := `
	INSERT INTO products(product_name, unit, category_id, price) VALUES ($1, $2, $3, $4)
	RETURNING product_id,product_name,category_id,unit,price ;
`
	err = db.Tx.QueryRow(query, product_name, unit, category_id, price).Scan(&product.Product_id, &product_name, &product.Category_id, &product.Unit, &product.Price)
	if err != nil {
		return product, err
	}

	query = `
	SELECT * FROM categories WHERE category_id = $1;
`
	err = db.Tx.QueryRow(query, product.Category_id).Scan(&product.Category.Category_id, &product.Category.Category_name, &product.Category.Description)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (db *DB) UpdateColumn(tableName, columnName string, idName string, tableId uint, value interface{}, product *Product) error {
	constQuery := fmt.Sprintf("UPDATE %s SET %s = $1 WHERE %s = %v RETURNING price ;", tableName, columnName, idName, tableId)
	err := db.Tx.QueryRow(constQuery, value).Scan(&product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ProductDelete(product_id uint) (err error) {
	query := `DELETE FROM products WHERE product_id = $1;`
	_, err = db.Tx.Exec(query, product_id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()
	db.Tx, err = db.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer db.Tx.Rollback()
	product, err := db.InsertIntoProduct("Gosht", "6.kg", 8, 10000.34)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product ID: ", product.Product_id)
	fmt.Println("Product: ", product.Product_name)
	fmt.Println("Category ID: ", product.Category_id)
	fmt.Println("Unit: ", product.Unit)
	fmt.Println("Description: ", product.Unit)
	fmt.Println("Price: ", product.Price)
	fmt.Println("------------------")
	fmt.Println("Category id: ", product.Category.Category_id)
	fmt.Println("Category name: ", product.Category.Category_name)
	fmt.Println("Description: ", product.Category.Description)

	err = db.UpdateColumn("products", "price", "product_id", product.Product_id, 50000, &product)
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------- ")
	fmt.Println("Price: ", product.Price)
	fmt.Println("Category id: ", product.Category.Category_id)
	fmt.Println("Category name: ", product.Category.Category_name)
	fmt.Println("Description: ", product.Category.Description)
	fmt.Println("-----------------")

	err = db.ProductDelete(product.Product_id)
	fmt.Println("Product successfully deleted! ")

	err = db.Tx.Commit()
	if err != nil {
		panic(err)
	}
}
