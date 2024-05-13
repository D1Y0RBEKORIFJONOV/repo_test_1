package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "+_+diyor2005+_+"
	dbname   = "demo"
)

type DB struct {
	db  *sql.Tx
	err error
}

func Connect() (*sql.DB, error) {
	InfoStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", InfoStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) CreateTableUsers() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ;
`
	_, db.err = db.db.Exec(query)
	if db.err != nil {
		return db.err
	}
	return nil
}

func (db *DB) CreateTablesFriend() error {
	query := `
	CREATE TABLE IF NOT EXISTS friendships (
    friendship_id SERIAL PRIMARY KEY,
    user_id INT,
    friend_id INT,
    status VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON SET DEFAULT
    FOREIGN KEY (friend_id) REFERENCES users(user_id)
);
`
	_, db.err = db.db.Exec(query)
	if db.err != nil {
		return db.err
	}
	return nil
}

func (db *DB) CreateUser(username, email, password string) (user User, err error) {
	query := `
	INSERT INTO users (username,email,password) VALUES ($1,$2,$3)
	RETURNING user_id,created_at ;
`
	err = db.db.QueryRow(query, username, email, password).Scan(&user.User_id, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	user.Email = email
	user.Password = password
	user.Username = username
	return user, nil
}

func (db *DB) InsertFriendship(user User, friend_id int) error {
	query := `
	INSERT INTO friendships(user_id, friend_id) VALUES ($1,$2)
`
	_, db.err = db.db.Exec(query, user.User_id, friend_id)
	if db.err != nil {
		return db.err
	}
	return nil
}

func (db *DB) SetStatus(user *User, friend_id int, status string) error {
	query := `
	UPDATE friendships 
	SET status = $1
	WHERE user_id = $2 AND friend_id = $3
`
	_, db.err = db.db.Exec(query, status, user.User_id, friend_id)
	if db.err != nil {
		return db.err
	}
	return nil
}

type User struct {
	User_id   int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func main() {
	datebase, err := Connect()
	if err != nil {
		panic(err)
	}
	defer datebase.Close()
	db := DB{}
	db.db, err = datebase.Begin()
	if err != nil {
		panic(err)
	}

	if db.err != nil {
		panic(db.err)
	}
	defer db.db.Rollback()
	err = db.CreateTableUsers()
	if err != nil {
		panic(err)
	}
	err = db.CreateTablesFriend()
	if err != nil {
		panic(err)
	}
	usr, err := db.CreateUser("Diyorbek", "diyordev3@gmail.com", "+_+diyor2005+_+")
	if err != nil {
		panic(err)
	}
	usr2, err := db.CreateUser("Kamron", "kamrondev4@gmail.com", "kama1234")
	if err != nil {
		panic(err)
	}

	err = db.InsertFriendship(usr, usr2.User_id)
	if err != nil {
		panic(err)
	}
	err = db.SetStatus(&usr, usr2.User_id, "request")
	if err != nil {
		panic(err)
	}
	fmt.Println("User status requested!")
	err = db.SetStatus(&usr, usr2.User_id, "accept")
	if err != nil {
		panic(err)
	}
	fmt.Println("User status accepted!")

	err = db.SetStatus(&usr, usr2.User_id, "blocked")
	if err != nil {
		panic(err)
	}
	fmt.Println("User status blocked!")

	err = db.db.Commit()
}
