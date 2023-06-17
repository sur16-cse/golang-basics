package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"` // key here will be lowercase after giving aliases
	Price int // key here will be  not lowercase there is no aliases
	Platform string `json:website` // key here will be lowercase after giving aliases
	Password string `json:"-"` // password hide it will not be show in json
	Tags []string `json:"tags, omitempty"`  // key here will be lowercase after giving aliases and it can be empty
}

func main()  {
	fmt.Println("Welcome to Json Video")
	EncodeJson()
}

func EncodeJson(){
	lcoCourses:=[]course{
		{"ReactJs Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"Mern Bootcamp",399,"LearnCodeOnline.in","abc1233",[]string{"full-stack","js"}},
		{"Angular Bootcamp",199,"LearnCodeOnline.in","abc12336",nil},
	}

	//package this data as json data
	finalJson,err:= json.Marshal(lcoCourses)
	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)

	endRes,err:=json.MarshalIndent(lcoCourses,"","\t")
	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s\n",endRes)
}