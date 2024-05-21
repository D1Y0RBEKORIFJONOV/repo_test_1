package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	dbname   = "student_course"
	user     = "postgres"
	password = "+_+diyor2005+_+"
)

func (db *DB) Connect() error {
	connectIfnoStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	var err error
	db.DB, err = sql.Open("postgres", connectIfnoStr)
	if err != nil {
		return err
	}

	err = db.DB.Ping()
	if err != nil {
		return nil
	}

	return nil
}

func (db *DB) InserIntoAllbum(id, title, artist string, price float64) (err error) {
	err = db.Connect()
	if err != nil {
		return err
	}
	defer db.DB.Close()
	insertIntoStr := `
	INSERT INTO album (id, title, artist,price) VALUES ($1, $2, $3,$4)
	RETURNING *;
	`

	_,err = db.DB.Exec(insertIntoStr, id, title, artist, price)

	if err != nil {
		return  err
	}

	return  err
}

func (db * DB)Delete(id string) error {
	err := db.Connect()
	if err != nil {
		return err
	}

	defer db.DB.Close()

	qery := `
	DELETE FROM album WHERE id = $1 ;
	`

	_,err  = db.DB.Exec(qery)

	if err != nil {
		return err
	}

	return nil
}