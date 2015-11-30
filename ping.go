package main

import (
	"log"

	"github.com/franela/goreq"
)

var targetURL = "http://localhost:1234"

func main() {
	for {
		doRequest()
	}
}

func doRequest() {
	res, err := goreq.Request{Uri: targetURL}.Do()
	if err != nil {
		log.Println(err)
	}
	if res != nil {
		log.Println(res.Status)
	}
}
