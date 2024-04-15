// 1) 12 chi uy vazifada ValidateUser funksiya yaratilgan edi.
// 2) Bugungi test_repo repository yaratilgan ichida o'zgarishlar kiritilishi kerak
// 3) ValidateUser dasturri test_repo ni ichiga joylang
// 4) user package yaratib uning ichida Validate digan funksiyani e'lon qiling
// 6) main package da user.Validate qilingan holda chaqirinsin
// 7) dastur ishlashini hosil qiling (go run yoki go build o'rqali)
// 8) barchasi ishlagan taqdirda, hamma o'zgarishlarni git add, git commit va git push commandar o'rqali bajaring
// 9) natijada local codelaringiz github ga joylangan bolishi kerak
// 10) Repository linkini jonating

package main

import (
	"fmt"

	validUser "github.com/D1Y0RBEKORIFJONOV/repo_test_1/14_uyga_vazifa/user"
	"github.com/D1Y0RBEKORIFJONOV/repo_test_1/14_uyga_vazifa/validateUser"
)

func main() {
	var usr  = validateUser.User{}
	validateUser.Input("Enter your name: ",&usr.Name)
	validateUser.Input("Enter your age: ",&usr.Age)
	validateUser.Input("Enter your email: ",&usr.Email)
	validErros := validUser.Validate(usr)
	if len(validErros) != 0 {
		for _,err := range validErros {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(usr.GetData())
}
