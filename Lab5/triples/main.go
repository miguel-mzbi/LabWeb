package main

import (
	"log"
	"time"

	triplestore "github.com/wallix/triplestore"
)

type Address struct {
	Street string `predicate:"street"`
	City   string `predicate:"city"`
}

type Person struct {
	Name     string    `predicate:"name"`
	Age      int       `predicate:"age"`
	Gender   string    `predicate:"gender"`
	Birth    time.Time `predicate:"birth"`
	Surnames []string  `predicate:"surnames"`
	Addr     Address   `predicate:"address" bnode:"myaddress"` // empty bnode value will make bnode value random
}

func main() {
	addr := &Address{Street: "Ramon Corona", City: "Guadalajara"}
	person := &Person{Name: "Miguel", Age: 21, Gender: "Male", Birth: time.Now(), Surnames: []string{"Montoya", "Zaragoza"}, Addr: *addr}

	tris := triplestore.TriplesFromStruct("Miguel", person)

	src := triplestore.NewSource()
	for _, triple := range tris {
		src.Add(triple)
	}
	snap := src.Snapshot()

	log.Println(snap.Contains(triplestore.SubjPred("Miguel", "name").StringLiteral("Miguel")))
	log.Println(snap.Contains(triplestore.SubjPred("Miguel", "age").IntegerLiteral(21)))
	log.Println(snap.Contains(triplestore.SubjPred("Miguel", "gender").StringLiteral("Male")))
	log.Println(snap.Contains(triplestore.SubjPred("Miguel", "address").Bnode("myaddress")))
	log.Println(snap.Contains(triplestore.BnodePred("myaddress", "street").StringLiteral("Ramon Corona")))
	log.Println(snap.Contains(triplestore.BnodePred("myaddress", "city").StringLiteral("New York")))
}
