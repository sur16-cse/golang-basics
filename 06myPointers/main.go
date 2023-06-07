package main

import "fmt"

func main()  {
	fmt.Println("Welcome to golang pointers")
	var ptr *int
	fmt.Println("value of pointer is,",ptr)

	myNumber := 12

	ptr=&myNumber
	fmt.Println("value of actual pointer is,",ptr)
	fmt.Println("value of pointer is,",*ptr)

	*ptr=*ptr+2;
	fmt.Println("New value of myNumber is,",myNumber)
}