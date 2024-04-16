package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name       string
	Age        int
	Occupation string
}

func GetUserFromReadFile(fileNmae string) (*[]User, error) {
	file, err := os.Open(fileNmae)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	usrs := []User{}
	for scanner.Scan() {
		tempUser := User{}
		for scanner.Text() != "" {
			usrArray := strings.Split(scanner.Text(), " ")
			if usrArray[0] == "Name:" {
				for i := 1; i < len(usrArray); i++ {
					tempUser.Name += usrArray[i] + " "
				}
			} else if usrArray[0] == "Age:" {
				tempUser.Age, _ = strconv.Atoi(usrArray[1])
			} else if usrArray[0] == "Occupation:" {
				for i := 1; i < len(usrArray); i++ {
					tempUser.Occupation += usrArray[i] + " "
				}
			}
			scanner.Scan()
		}
		usrs = append(usrs, tempUser)
	}

	return &usrs, nil
}

func PrintUserToFile(fileName string, usr User) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	constStr := fmt.Sprintf("\n\nName: %s\nAge: %d\nOccupation: %s", usr.Name, usr.Age, usr.Occupation)
	_, err = file.WriteString(constStr)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	command := 0
	fmt.Print("1.ReadFile\t2.AddUser\n0.help\t\t-1.exit\n")
	for {
		fmt.Print("Enter command[0.help]: ")
		fmt.Scan(&command)
		switch command {
		case -1:
			return
		case 0:
			fmt.Print("1.ReadFile\t\t2.AddUser\n0.help\t\t-1.exit\n")
		case 1:
			usrs, err := GetUserFromReadFile("/home/diyorbek/go/src/repo_test_1/15_uyga_vazifa/sample.txt")
			if err != nil {
				log.Fatal(err)
			}
			for _, usr := range *usrs {
				fmt.Printf("Name: %s\nAge: %d\nOccupation: %s\n\n", usr.Name, usr.Age, usr.Occupation)
			}
		case 2:
			user := User{}
			fmt.Print("Name: ")
			fmt.Scan(&user.Name)
			fmt.Print("Age: ")
			fmt.Scan(&user.Age)
			fmt.Print("Occupation: ")
			fmt.Scan(&user.Occupation)
			_ = PrintUserToFile("/home/diyorbek/go/src/repo_test_1/15_uyga_vazifa/sample.txt", user)
		}
	}
}
