package main

import "fmt"

func main()  {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	fmt.Println("Hello")

	myDefer()
}

//defer make thing executable from lifo order

func myDefer()  {
	for i:=0;i<5;i++{
	  defer	fmt.Print(i)
	}
}