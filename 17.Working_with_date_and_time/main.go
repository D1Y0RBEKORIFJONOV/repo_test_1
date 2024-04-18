package main

import (
	"encoding/json"
	"os"
)

// Define a struct that matches the structure of your JSON data
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

func main() {
	file, err := os.Open("employees.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		panic(err)
	}

	usr := User{Name: "Diyorbek", Age: 19, ID: 6, Position: "Goland ProgrAmer"}
	users = append(users, usr)
	file, err = os.OpenFile("employeesCopy.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(jsonData)
}
