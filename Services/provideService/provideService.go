package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Prize : ...
type Prize struct {
	Category string `json:"category"`
}

// Winner : ...
type Winner struct {
	ID            string  `json:"id"`
	Firstname     string  `json:"firstname"`
	Lastname      string  `json:"surname"`
	PrizeCategory []Prize `json:"prizes"`
}

// Laureates : ...
type Laureates struct {
	Winners []Winner `json:"Laureates"`
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// winner1 := Winner{"1", "winner1Name", "winner1Lastname"}
	// winner2 := Winner{"2", "winner2Name", "winner2Lastname"}
	// winners := []Winner{winner1, winner2}
	// laureates := Laureates{winners}

	laureates := Laureates{
		[]Winner{
			Winner{"2", "winner2Name", "winner2Lastname", []Prize{
				Prize{"death"}}},
			Winner{"1", "winner1Name", "winner1Lastname", []Prize{
				Prize{"life"},
				Prize{"eating"},
			}},
		},
	}

	js, _ := json.Marshal(laureates)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe("localhost:1337", nil)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
		return
	}
}
