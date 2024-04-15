package validUser

import "github.com/D1Y0RBEKORIFJONOV/repo_test_1/14_uyga_vazifa/validateUser"


func Validate(usr validateUser.User) []error {
	var validErros []error
	err := usr.CheckName()
	if err != nil {
		validErros = append(validErros, err)
	}
	err = usr.CheckEmail()
	if err != nil {
		validErros = append(validErros, err)
	}

	err = usr.CheckAge() 
	if  err != nil{
		validErros = append(validErros, err)
	}

	if len(validErros) != 0 {
		return validErros
	}

	return nil
}