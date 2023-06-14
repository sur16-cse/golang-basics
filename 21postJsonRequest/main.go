package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcomne to getRequest golang")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/post"
	requestBody := strings.NewReader(`{"coursename":"Let's Go with GoLang","price":0,"platfoem":"learnCodeOnline.in"} `)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response body:", string(content))
}
