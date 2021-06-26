package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Url_parser_withinterface() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var f interface{}

	err = json.Unmarshal(body, &f)

	fmt.Printf("%+v\n", f)
}
