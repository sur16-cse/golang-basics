package main

import (
	"fmt"

	
)

func main()  {
	fmt.Println("Welcome to golang maps")
	languages := make(map[string]string)

	languages["JS"]="JavaScript"
	languages["RB"]="Ruby"
	languages["PY"]="Python"

	fmt.Println("list of all languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])

	delete(languages,"RB")
	fmt.Println("list of all languages: ",languages)

	//loops are interesting in golang
	for key,value := range languages{
		fmt.Printf("For key %v,value is %v\n",key,value)
	}

	for _,value := range languages{
		fmt.Printf("For key v is,value is %v\n",value)
	}
}