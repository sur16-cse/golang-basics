package main

import "fmt"

func main()  {
	fmt.Println("Welcome to structs")
	// no inheritance in golang; no super or parent

	Surbhi := User{"Surbhi","Surbhi@go.dev",true,16}
	fmt.Println(Surbhi)
	fmt.Printf("Surbhi details are: %+v\n",Surbhi )
	fmt.Printf("Name is %v and email is %v",Surbhi.Name,Surbhi.Email)
	fmt.Println("")
	Surbhi.GetStatus()
	Surbhi.NewEmail()
	fmt.Printf("Name is %v and email is %v",Surbhi.Name,Surbhi.Email)
	
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus()  {
	fmt.Println("Is user active",u.Status)
}

func (u User) NewEmail()  {
	u.Email="test@go.dev"
	fmt.Println("Email of the user is",u.Email)
}