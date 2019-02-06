package main

import (
	"encoding/json"
	"gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

// Item : ...
type Item struct {
	ID    int    `json:"Id"`
	Value string `json:"value"`
}

// Items : ...
type Items struct {
	Items []Item `json:"Items"`
}

// Message : ...
type Message struct {
	Message string `json:"message"`
}

var database = Items{[]Item{}}

// Appends item to database
func (database *Items) addItemToDB(item Item) {
	database.Items = append(database.Items, item)
}

func getItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	for _, item := range database.Items {
		if item.ID == id {
			js, _ := json.Marshal(item)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}
	}
	js, _ := json.Marshal(Message{"OBJECT NOT FOUND"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getItems(w http.ResponseWriter, req *http.Request) {
	js, _ := json.Marshal(database)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func deleteItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	for i, item := range database.Items {
		if item.ID == id {
			js, _ := json.Marshal(item)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

			database.Items = database.Items[:i+copy(database.Items[i:], database.Items[i+1:])]
			return
		}
	}
	js, _ := json.Marshal(Message{"OBJECT NOT FOUND"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func deleteItems(w http.ResponseWriter, req *http.Request) {
	database.Items = []Item{}
	js, _ := json.Marshal(Message{"ITEMS DELETED"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func editItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	for i, item := range database.Items {
		if item.ID == id {
			newValue, _ := req.URL.Query()["newValue"]
			database.Items[i].Value = newValue[0]
			js, _ := json.Marshal(database.Items[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}
	}
	js, _ := json.Marshal(Message{"OBJECT NOT FOUND"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func addItem(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var item Item
	err := decoder.Decode(&item)
	if err != nil {
		log.Fatal(err)
	}
	database.addItemToDB(item)
	js, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func indexHandler(entrypoint string) func(w http.ResponseWriter, req *http.Request) {
	fn := func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, entrypoint)
	}
	return http.HandlerFunc(fn)
}

func main() {
	mxRouter := mux.NewRouter()
	api := mxRouter.PathPrefix("/api").Subrouter()
	api.HandleFunc("/getItems/{id}", getItem).
		Methods("GET")
	api.HandleFunc("/getItems", getItems).
		Methods("GET")
	api.HandleFunc("/deleteItem/{id}", deleteItem).
		Methods("DELETE")
	api.HandleFunc("/deleteItems", deleteItems).
		Methods("DELETE")
	api.HandleFunc("/editItem/{id}", editItem).
		Methods("PUT")
	api.HandleFunc("/addItem", addItem).
		Methods("POST")

	mxRouter.PathPrefix("/static").Handler(http.FileServer(http.Dir("dist/")))
	mxRouter.PathPrefix("/").HandlerFunc(indexHandler("dist/index.html"))

	corsRouter := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:1337"},
		AllowCredentials: true,
	})

	handler := corsRouter.Handler(mxRouter)

	err := http.ListenAndServe(":1337", handler)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
		return
	}
}
