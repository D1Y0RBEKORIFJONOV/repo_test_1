package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	dbname   = "store"
	user     = "postgres"
	port     = 5432
	password = "+_+diyor2005+_+"
)

type DB struct {
	DB *sql.DB
	Tx *sql.Tx
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

func (db *DB) CreateTableUsers() error {
	query := `CREATE TABLE  IF NOT EXISTS users (user_id SERIAL PRIMARY KEY,username VARCHAR(255),balance FLOAT);`
	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) CreateTableProducts() error {
	query := `CREATE TABLE IF NOT EXISTS products(product_id SERIAL PRIMARY KEY,name VARCHAR(50),count INT,price FLOAT);`
	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) InsertToAutoUsers(name string, balance float64) error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.DB.Close()

	query := `INSERT INTO users(username,balance) VALUES ($1,$2);`
	_, err = db.DB.Exec(query, name, balance)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) CreateTableProductUsers() error {
	query := `CREATE TABLE IF NOT EXISTS product_users (product_id SERIAL PRIMARY KEY,name VARCHAR(50),user_id INT REFERENCES users(user_id));`
	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Migration() error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.DB.Close()
	err = db.CreateTableUsers()
	if err != nil {
		fmt.Println("1")
		return err
	}
	err = db.CreateTableProducts()
	if err != nil {
		fmt.Println("2")
		return err
	}
	err = db.CreateTableProductUsers()
	if err != nil {
		fmt.Println("3")
		return err
	}
	return nil
}

func (db *DB) InsertToUser(username string, balance float64) (user User, err error) {
	db, err = ConnectDb()
	if err != nil {
		return User{}, err
	}
	defer db.DB.Close()
	query := `INSERT INTO users(username,balance) VALUES ($1,$2) RETURNING user_id;`

	err = db.DB.QueryRow(query, username, balance).Scan(&user.User_id)
	if err != nil {
		return User{}, err
	}
	user.Username = username
	user.Balance = balance
	return user, nil
}

func (db *DB) InsertToProducts(name string, count int, price float64) (products Products, err error) {
	db, err = ConnectDb()
	if err != nil {
		return Products{}, err
	}
	defer db.DB.Close()
	query := `INSERT INTO products(name,count,price) VALUES ($1,$2,$3) RETURNING product_id;`
	err = db.DB.QueryRow(query, name, count, price).Scan(&products.Product_id)
	if err != nil {
		return Products{}, err
	}
	products.Name = name
	products.Count = count
	products.Price = price
	return products, nil
}

func (db *DB) InsertProduct(name string, user_id int) (product Product, err error) {
	query := `INSERT INTO product_users(name,user_id) VALUES ($1,$2) RETURNING product_id;`
	err = db.Tx.QueryRow(query, name, user_id).Scan(&product.Product_id)
	if err != nil {
		return Product{}, err
	}
	product.Name = name
	product.User_id = user_id
	return product, nil
}

func (db *DB) UpdateCount(products *Products) error {
	query := `UPDATE products SET count = count - 1 WHERE product_id = $1 RETURNING count;`
	err := db.Tx.QueryRow(query, products.Product_id).Scan(&products.Count)
	if err != nil {
		return err
	}
	if products.Count < 0 {
		return errors.New("Product count is negative")
	}
	return nil
}

func (db *DB) GetProduct(user *User, product *Products) error {
	query := `UPDATE users SET balance = balance - $1 WHERE user_id  = $2 RETURNING balance;`

	err := db.Tx.QueryRow(query, product.Price, user.User_id).Scan(&user.Balance)
	if err != nil {
		return err
	}
	if user.Balance < 0 {
		return errors.New("User balance is negative")
	}
	return nil

}

type User struct {
	User_id  int
	Username string
	Balance  float64
	Product  []Product
}
type Product struct {
	Product_id int
	Name       string
	User_id    int
}

type Products struct {
	Product_id int
	Name       string
	Count      int
	Price      float64
}

func main() {
	var (
		err error
		db1 DB
	)
	err = db1.Migration()
	if err != nil {
		panic(err)
	}
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()
	users, err := db.InsertToUser("Kamronbek", 3.34)
	if err != nil {
		panic(err)
	}
	product, err := db.InsertToProducts("Hurmo", 0, 65.34)
	if err != nil {
		panic(err)
	}
	fmt.Println("User:")
	fmt.Println("users Name:", users.Username)
	fmt.Println("User balance:", users.Balance)
	fmt.Println("User product", users.Product)

	fmt.Println("Product: ")
	fmt.Println("product Name:", product.Name)
	fmt.Println("product Price:", product.Price)
	fmt.Println("product Count:", product.Count)

	db.Tx, err = db.DB.Begin()
	if err != nil {
		panic(err)
	}
	err = db.UpdateCount(&product)
	if err != nil {
		panic(err)
	}
	err = db.GetProduct(&users, &product)
	if err != nil {
		panic(err)
	}

	prr, err := db.InsertProduct(product.Name, users.User_id)
	if err != nil {
		panic(err)
	}

	users.Product = append(users.Product, prr)

	err = db.Tx.Commit()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nUser:")
	fmt.Println("users Name:", users.Username)
	fmt.Println("User balance:", users.Balance)
	fmt.Println("User product", users.Product)

	fmt.Println("Product: ")
	fmt.Println("product Name:", product.Name)
	fmt.Println("product Price:", product.Price)
	fmt.Println("product Count:", product.Count)

}
