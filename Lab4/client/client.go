package main

import (
	"log"
	"net/http"
)

func readCookie() {
	resp, err := http.Get("http://localhost:1337/cookie")
	if err != nil {
		log.Fatal("Error reading cookie: ", err)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "Hi" {
			log.Println("Cookie: ", cookie.Value)
		}
	}
}

func main() {
	readCookie()
}
