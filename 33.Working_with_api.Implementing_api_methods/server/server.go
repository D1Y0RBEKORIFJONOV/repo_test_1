package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	lisiner, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lisiner.Close()
	fmt.Println("Server is listening on port 8080:")
	for {
		conn, err := lisiner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		sourse := string(buf[0:n])
		file, err := os.Open(sourse)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write([]byte("File mufaqiyatli yuborildi!"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("adding file to connection", file.Name())
		defer file.Close()

	}
}
