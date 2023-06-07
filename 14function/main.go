package main

import "fmt"

func main()  {
	fmt.Println("Welcome to function in golang")
	greeter()

	//cannot create function inside functtion
	// func greeterTwo()  {
	// 	fmt.Println("Another Method")
	// }

	result:=adder(3,5)

	fmt.Println("result is, ", result)

	prores:=proadder(4,5,7,9,3)
	fmt.Println("prores is",prores)

	proressub,mymessage:=prosub(10,1,2,3,1)
	fmt.Println("proresub is ",proressub," ","promessage is ",mymessage)
}

func adder(a int,b int) int {
	return a+b
}

func proadder(values ...int) int  {
	total:=0

	for _,val:=range values{
		total+=val
	}

	return total;
}

func prosub(values ...int) (int,string)  {
	total:=0

	for _,val:=range values{
		total-=val
	}

	return total,"Hi i am substract";
}


func greeter()  {
	fmt.Println("Namastey from golang")
}