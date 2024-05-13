package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		var sourse string = ""
		_, err := fmt.Scanln(&sourse)
		if err != nil {
			panic(err)
		}
		if n, err := conn.Write([]byte(sourse)); err != nil || n == 0 {
			fmt.Println("Write Error")
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read Error")
		}
		fmt.Println(string(buf[0:n]))
		fmt.Println()
	}
}
