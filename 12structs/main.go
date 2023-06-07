package main

import "fmt"

func main()  {
	fmt.Println("Welcome to structs")
	// no inheritance in golang; no super or parent

	Surbhi := User{"Surbhi","Surbhi@go.dev",true,16}
	fmt.Println(Surbhi)
	fmt.Printf("hitesh details are: %+v\n",Surbhi )
	fmt.Printf("Name is %v and email is %v",Surbhi.Name,Surbhi.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}