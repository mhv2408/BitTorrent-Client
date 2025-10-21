package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func request(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("cannot get the url: ", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading data from response: ", err)
	}
	fmt.Println("Response Body: ")
	return string(data)
}
