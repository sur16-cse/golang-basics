package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"` // key here will be lowercase after giving aliases
	Price    int      // key here will be  not lowercase there is no aliases
	Platform string   `json:"Platform"`        // key here will be lowercase after giving aliases
	Password string   `json:"-"`               // password hide it will not be show in json
	Tags     []string `json:"tags, omitempty"` // key here will be lowercase after giving aliases and it can be empty
}

func main() {
	fmt.Println("Welcome to Json Video")
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`[{
                "coursename": "ReactJs Bootcamp",
                "Price": 299,
                "Platform": "LearnCodeOnline.in",
                "tags": [
                        "web-dev",
                        "js"
                ]
        },
        {
                "coursename": "Mern Bootcamp",
                "Price": 399,
                "Platform": "LearnCodeOnline.in",
                "tags": [
                        "full-stack",
                        "js"
                ]
        },
        {
                "coursename": "Angular Bootcamp",
                "Price": 199,
                "Platform": "LearnCodeOnline.in",
                "tags": null
        }]`)

	var lcoCourses []course //slice of course struct because we have array of courses in json
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("Json is not valid")
	}

	//some cases you just want to add data to key value
	var myOnlineData []map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for _, v := range myOnlineData {
		// fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
		for i, value := range v {
			fmt.Printf("Key is %v and value is %v and Type is %T\n", i, value, value)
		}
		fmt.Print("\n")
	}
}
