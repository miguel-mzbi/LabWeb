package main

import (
	"log"
	"net/url"

	"github.com/piprate/json-gold/ld"
)

func main() {
	serviceURL := "https://kgsearch.googleapis.com/v1/entities:search"

	params := url.Values{}
	params.Add("query", "Taylor Swift")
	params.Add("limit", "50")
	params.Add("indent", "true")
	params.Add("key", "INSERT API KEY HERE"
	par := params.Encode()

	addr := serviceURL + "?" + par

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	expanded, err := proc.Expand(addr, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}

	all := expanded[0].(map[string]interface{})["http://schema.org/itemListElement"]
	elementList := all.([]interface{})
	element := elementList[len(elementList)-1]
	content := element.(map[string]interface{})["http://schema.org/result"]
	descWrapper := content.([]interface{})[0].(map[string]interface{})["http://schema.org/description"]
	descValue := descWrapper.([]interface{})[0].(map[string]interface{})["@value"]

	// fmt.Printf("%# v", pretty.Formatter(all))
	log.Println(descValue)
}
