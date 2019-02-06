// Miguel Angel Montoya
// A01226045
package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func f1(ch chan int) {
	ch <- 100
}

func f2(ch chan int) {
	i := <-ch
	log.Println(i)
}

func f4() {
	ch := make(chan int)
	go f1(ch)
	go f2(ch)
	time.Sleep(50 * time.Millisecond)
}

func getHeaders(site string, ch chan string, waitGroup *sync.WaitGroup) {
	res, err := http.Get(site)
	if err != nil {
		log.Panicln(err)
		waitGroup.Done()
	} else {
		dateString := res.Header.Get("Date")
		waitGroup.Done()
		ch <- string(dateString)
		// log.Println("HEADER", "\n")
	}
}

func iterateSites() {
	sites := [5]string{
		"http://google.com",
		"http://yahoo.com",
		"http://mitec.itesm.mx",
		"http://miscursos.tec.mx",
		"http://example.com"}

	ch := make(chan string)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(sites))

	for _, site := range sites {
		go getHeaders(site, ch, &waitGroup)
	}

	waitGroup.Wait()

	for i := 0; i < len(sites); i++ {
		element := <-ch
		log.Println(element)
	}
}

func main() {
	iterateSites()
}
