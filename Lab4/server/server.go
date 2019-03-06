package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type People struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender"`
	Age       string   `json:"age"`
	EyeColor  string   `json:"eye_color"`
	HairColor string   `json:"hair_color"`
	Films     []string `json:"films"`
	Species   string   `json:"species"`
	URL       string   `json:"url"`
}

type PeopleCont struct {
	People []People
	Error  error
}

func cookieHandler(w http.ResponseWriter, req *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "Hi", Value: "Miguel & Arturo", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func cookieReader(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Hi")
	if err != nil {
		log.Fatal("Error reading cookie: ", err)
	} else {
		value := cookie.Value
		log.Println("The cookie value: ", value)
	}
}

func readJSON(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	id := ""
	if q["id"] != nil {
		id = q["id"][0]
	}

	c := make(chan PeopleCont)
	go fetchChanJSON(id, c)

	var peopleC = <-c
	if peopleC.Error != nil {
		log.Fatalln("Error on reading JSON: ", peopleC.Error)
	}

	bytes, err := json.Marshal(peopleC.People)

	// var peopleC = fetchJSON(id)
	// bytes, err := json.Marshal(peopleC)

	if err != nil {
		log.Fatal("Error marshal: ", err)
		w.WriteHeader(500)
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(bytes)
	}
}

func fetchJSON(id string) []People {
	log.Println("Non-concurrent")
	resp, err := http.Get("https://ghibliapi.herokuapp.com/people/" + id)
	if err != nil {
		log.Fatal("Error reading Ghibli: ", err)
		return []People{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body: ", err)
		return []People{}
	}

	if id == "" {
		var people []People
		json.Unmarshal(body, &people)
		return people
	} else {
		var people People
		json.Unmarshal(body, &people)
		return []People{people}
	}
}

func fetchChanJSON(id string, c chan PeopleCont) {
	log.Println("Concurrent")
	resp, err := http.Get("https://ghibliapi.herokuapp.com/people/" + id)
	if err != nil {
		c <- PeopleCont{People: []People{}, Error: errors.New("Error reading Ghibli")}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- PeopleCont{People: []People{}, Error: errors.New("Error reading body")}
	}

	if id == "" {
		var people []People
		json.Unmarshal(body, &people)
		c <- PeopleCont{People: people, Error: nil}
	} else {
		var people People
		json.Unmarshal(body, &people)
		c <- PeopleCont{People: []People{people}, Error: nil}
	}
}

func main() {
	http.HandleFunc("/cookie", cookieHandler)
	http.HandleFunc("/readCookie", cookieReader)
	http.HandleFunc("/readJSON", readJSON)

	err := http.ListenAndServe("localhost:1337", nil)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
		return
	}

}
