package main

import (
	"log"
	"sync"
)

func f1(w *sync.WaitGroup) {
	log.Println("This is F1")
	w.Done()
}

func f2(w *sync.WaitGroup) {
	log.Println("This is F2")
	w.Done()
}

func f3(w *sync.WaitGroup) {
	log.Println("This is F3")
	w.Done()
}

func f4(w *sync.WaitGroup) {
	go f1(w)
	go f2(w)
	go f3(w)
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	f4(&waitGroup)
	waitGroup.Wait()
}
