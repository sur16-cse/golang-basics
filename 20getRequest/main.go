package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcomne to getRequest golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Response status:", response.StatusCode)
	fmt.Println("content length:", response.ContentLength)

	var responseString strings.Builder
	content,_:= ioutil.ReadAll(response.Body)
	byteCount,_:= responseString.Write(content)
	fmt.Println("byteCount:", byteCount)
	fmt.Println("Response body:", responseString.String())
	
	fmt.Println("Response body:", string(content))
}
