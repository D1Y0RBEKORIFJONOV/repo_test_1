package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "+_+diyor2005+_+"
	dbname   = "n9_go"
)

type DB struct {
	DB *sql.DB
}

func Connect() (*sql.DB, error) {
	info := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", info)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) ShowColumnWithExplain(columnName string, value interface{}) error {
	query := `
	EXPLAIN ANALYZE SELECT * FROM people WHERE $1 =  $2 ;
`
	rows, err := db.DB.Query(query, columnName, value)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var explain string
		err = rows.Scan(&explain)
		if err != nil {
			return err
		}
		fmt.Println(explain)
	}
	return nil
}

func (db *DB) CreateIndex(indexName, columnName string) error {
	query := `
	CREATE INDEX $1 ON people ($1);
`
	_, err := db.DB.Exec(query, indexName, columnName)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) ShowLastAndFirstName(first_name, last_name string) error {
	query := `
	EXPLAIN ANALYZE SELECT * FROM people WHERE first_name =  $1 AND last_name =  $2 ;
`
	rows, err := db.DB.Query(query, first_name, last_name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var explain string
		err = rows.Scan(&explain)
		if err != nil {
			return err
		}
		fmt.Println(explain)
	}
	return nil
}
func (db *DB) DropIndex(indexName string) error {
	quer := `DROP INDEX $1 ;`
	_, err := db.DB.Exec(quer, indexName)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	db := DB{}
	var err error
	db.DB, err = Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("######################################")
	err = db.ShowColumnWithExplain("first_name", "Adrian")
	fmt.Println("######################################")
	err = db.CreateIndex("index_first_name", "first_name")

	fmt.Println("######################################")
	err = db.ShowColumnWithExplain("first_name", "Adrian")
	fmt.Println("######################################")

	// birinchi amaldan ikkinchisini Natijasi o'zgarishi:first_name ustuniga indexyaratishi sabali zapros tezroq bajarildi!!

	err = db.DropIndex("index_first_name")
	fmt.Println("droping index first_name")
	fmt.Println("######################################")

	_, err = db.DB.Exec("CREATE INDEX IF NOT EXISTS index_firstame_last_name ON people (first_name,last_name);")
	if err != nil {
		panic(err)
	}
	err = db.ShowLastAndFirstName("Adrian", "Gross")
	if err != nil {
		panic(err)
	}
	fmt.Println("#####################################")

	err = db.ShowLastAndFirstName("Gross", "Adrian")
	if err != nil {
		panic(err)
	}

	// Natija o'zgaradi chunki qiduruv tartibini o'zgatirish indexni foydalanishi oz'artiradi va tezlashtirai!!

}
