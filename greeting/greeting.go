package greeting

import (
	"github.com/D1Y0RBEKORIFJONOV/repo_test_1/14_uyga_vazifa/validateUser"
)

func Greet() string{
	return "Hi,How are you? "
}

func InputName()string {
	var name string
	validateUser.Input("ENTER YOUR NAME: ",&name)
	return name
}