package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data []struct {
	UserID string `json:"userId"`
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Url_parser_withstruct() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var f Data

	err = json.Unmarshal(body, &f)

	fmt.Printf("%+v", f[0].Body)
}
