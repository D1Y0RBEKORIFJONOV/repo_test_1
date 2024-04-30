package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

// WRITE DAGI DEADLINE EROR NI KORISH UCHUN time.sleep(time.Second * 4) commendan oling!!

func WriteFile(ctx context.Context, filename, data string) error {
	//time.Sleep(time.Second * 4)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, err = file.WriteString("\n" + data)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := WriteFile(ctx, "readAndWrite_uy_vazifasi/file.txt", "hello world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data successfully written")
}
