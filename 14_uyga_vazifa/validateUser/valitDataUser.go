//     kiritishini ValidateUser digan method yarating.
//
// Qoidalar:
// Name empty bo'lishi kerak emas
// Name uzunligi kamida 6 belgigi bo'lishi kerak
// Age 0 dan kotta va 120 dan kichik bo'lishi
// Email empty bo'lishi kerak emas
// Email formatiga mos bolishi kerak (masalan example@domain.com)
//
// 2. Error slice yaratilgan holda barcha paydo bo'lgan errorlarni qoshib yuvorin
// 3. Foydalanuvchi ma'lumotlarni terminaldan oqib oling
// 4. Oqib oliniyatgan jarayonda errorlarni ohirida chiqarib berin
//
// Masalan:
// Name: asd
// Age: 123
// Email: ""

// Errors:
// Name: length cannot be less than a 6 characters
// Age: couldn't be more than 120
// Email: couldn't be empty

package validateUser

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func Input(str string, obj interface{}) interface{} {
	fmt.Print(str)
	typ := reflect.TypeOf(obj)

	if typ.Kind() != reflect.Ptr {
		fmt.Println("Error: Object must be a pointer!!!")
		os.Exit(1)
	}

	newObj := reflect.New(typ.Elem()).Interface()

	_, err := fmt.Scan(newObj)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(newObj).Elem())

	return obj
}

type User struct {
	Name  string
	Age   int
	Email string
}

func Init(name string, age int, email string) User {
	return User{name, age, email}
}

func (u *User) IsEmpty(str, typeEmpty string) error {
	if len(str) != 0 {
		err := fmt.Sprintf("%s: couldn't be empty", typeEmpty)
		return errors.New(err)
	}
	return nil
}

func (u *User) CheckName() error {
	if !(len(u.Name) >= 6) {
		return errors.New("Name must be length 6!!")
	}
	if u.Email == "" {
		return errors.New("Email could been empty")
	}

	return nil
}

func (u *User) CheckAge() error {
	if !(u.Age > 0 && u.Age < 120) {
		return errors.New("Age Couldn't be more than 120!")
	}
	return nil
}
func (u *User) CheckEmail() error {
	if u.Email == "" {
		return errors.New("Email could been empty")
	}

	var isRight, isPoint bool = false, false
	str := ""
	for i := 0; i < len(u.Email); i++ {
		if u.Email[i] == '@' {
			isRight = true
		}
		if isRight && u.Email[i] == '.' {
			isPoint = true
		}
		if isPoint && isRight {
			str += string(u.Email[i])
		}
	}
	if str != ".com" {
		return errors.New("Email entered uncorrect!!")
	}

	return nil
}

func (u *User) GetData() string {
	return fmt.Sprintf("Name: %s,Age:%d,Email: %s", u.Name, u.Age, u.Email)
}

