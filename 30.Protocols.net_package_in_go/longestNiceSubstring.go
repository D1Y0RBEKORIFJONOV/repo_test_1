package main

import "fmt"

func longestNiceSubstring(s string) string {
	str := string(s[0])
	c := 1
	for i := 0; i < len(s); i++ {
		if str[c-1] == s[i]+32 || str[c-1] == s[i]-32 {
			str += string(s[i])
			c++
		} else if c >= 2 {
			return str
		}
	}
	if c >= 2 {
		return str
	}

	return ""
}

func main() {
	fmt.Println(longestNiceSubstring("YazaAay"))
}
