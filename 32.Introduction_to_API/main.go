package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"sync"
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

var ErrGroup errgroup.Group

func connect() *sql.DB {
	strInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", strInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
func InsertToTable(ctx context.Context, db *sql.DB, goErr *errgroup.Group, generated int) error {
	goErr.Go(func() error {
		query := "INSERT INTO large_dataset(generated) VALUES ($1) "
		_, err := db.Exec(query, generated)
		if err != nil {
			return err
		}
		return nil
	})
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return goErr.Wait()
	}
}

func SelectFromTable(ctx context.Context, db *sql.DB, goErr *errgroup.Group, id_tale int) (genered Genered, err error) {
	goErr.Go(func() error {
		query := "SELECT * FROM large_dataset WHERE id =$1"
		err = db.QueryRow(query, id_tale).Scan(&genered.ID, &genered.Genered)
		if err != nil {
			return err
		}
		return nil
	})
	select {
	case <-ctx.Done():
		return genered, ctx.Err()
	default:
		return genered, goErr.Wait()
	}

}

func UpdataTable(ctx context.Context, db *sql.DB, goErr *errgroup.Group, columnName, value interface{}, id int) error {
	goErr.Go(func() error {
		query := fmt.Sprintf("UPDATE large_dataset SET %s = %v WHERE id = %d ", columnName, value, id)
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
		return nil
	})
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return goErr.Wait()
	}
}

func main() {
	db := connect()
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	genered := Genered{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := InsertToTable(ctx, db, &ErrGroup, 1234)
		if err != nil {
			panic(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		genered, err = SelectFromTable(ctx, db, &ErrGroup, 1)
		if err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := UpdataTable(ctx, db, &ErrGroup, "generated", 345345, 1)
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()

	fmt.Println(genered)
}
