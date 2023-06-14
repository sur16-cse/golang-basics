package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcomne to getRequest golang")
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postform"
	
	//formdata
	data:=url.Values{}
	data.Add("firstname","surbhi")
	data.Add("lastname","kumari")
	data.Add("email","surbhikumaridav@gmail.com")
	response,err:=http.PostForm(myUrl,data)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
