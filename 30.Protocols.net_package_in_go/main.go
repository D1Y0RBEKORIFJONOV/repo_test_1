package main

import (
	"fmt"
	"log"
	"net/url"
)

func PrintDataUrl(url *url.URL) {
	fmt.Println("URL:", url.String())
	fmt.Println("Protocol:", url.Scheme)
	fmt.Println("Host:", url.Host)
	fmt.Println("Path:", url.Path)
	fmt.Println("Query:", url.RawQuery)
	fmt.Println("Fragment:", url.Fragment)
}

func isCorrectUrl(url *url.URL) bool {
	if url.Scheme != "" && url.Host != "" {
		return true
	}
	return false
}

// Hato url korish uchun commentdan oling!
func main() {
	var URL string = "http//erp.student.najottalim.uz/my-groups/1178/staff/1178"
	//URL = "https:/ww.example.com/path/to/pagefoo=bar&ba=qux#fragment"
	usl, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	if !isCorrectUrl(usl) {
		log.Fatal("Invalid URL")
	}

	PrintDataUrl(usl)
}
