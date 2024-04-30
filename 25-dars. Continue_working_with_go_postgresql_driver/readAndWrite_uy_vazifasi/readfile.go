package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

// READFILE DAGI DEADLINE EROR NI KORISH UCHUN time.sleep(time.Second * 4) commendan oling!!

func ReadFile(ctx context.Context, fileName string, data chan<- string) error {
	//time.Sleep(time.Second * 4)
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			data <- scanner.Text()
		}
	}

	close(data)
	return scanner.Err()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	data := make(chan string)
	go func() {
		err := ReadFile(ctx, "readAndWrite_uy_vazifasi/file.txt", data)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for str := range data {
		fmt.Println(str)
	}
}
