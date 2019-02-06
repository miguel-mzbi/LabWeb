package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Profile : Testing struct
type Profile struct {
	Name    string
	Hobbies []string
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "where", Value: "Miguel At /", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "where", Value: "Miguel At /test", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func test2Handler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/test", 300)
}

func exploreReqHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprintf(w, "Requested: %s", req.URL)
	// log.Println(req.Header)
}

func parseReqHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintln(w, req.Form)
}

func returnJSONHandler(w http.ResponseWriter, req *http.Request) {
	profile := Profile{"Miguel", []string{"cook", "code"}}
	js, _ := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/test2", test2Handler)
	http.HandleFunc("/exploreReq", exploreReqHandler)
	http.HandleFunc("/parseReq", parseReqHandler)
	http.HandleFunc("/returnJSON", returnJSONHandler)

	err := http.ListenAndServe("localhost:1337", nil)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
		return
	}
}
