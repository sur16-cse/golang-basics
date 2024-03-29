package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals=[]string{"test"}

var wg sync.WaitGroup//pointer
var mut sync.Mutex //pointer


func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals=append(signals, endPoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)
	}
}
