package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Laboratory #1")
}

func sheepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "There you go, 101 sheep.\nThanks for calling %s", r.URL.Path)

	t := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "\n\nTime of order %s", t)
}

func serveSheep(w http.ResponseWriter, r *http.Request) {
	name := "./index.html"
	http.ServeFile(w, r, name)

}

func main() {
	http.HandleFunc("/", darMensaje)
	http.HandleFunc("/wantSheep", sheepHandler)
	http.HandleFunc("/sheepPhoto", serveSheep)

	err := http.ListenAndServe("localhost"+":"+"1337", nil)
	if err != nil {
		log.Fatal("error en el servidor : ", err)
		return
	}
}
