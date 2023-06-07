package main

import "fmt"

func main()  {
	fmt.Println("Welcome to array in golang")

	var fruitList [5]string
	fruitList[0]="Mango"
	fruitList[1]="Apple"
	fruitList[2]="Tomato"
	fruitList[4]="Peach"

	fmt.Println("Fruit List is: ",fruitList)
	fmt.Println("Fruit List is: ",len(fruitList))

	var vegList=[3]string{"potato","beans","mushroom"}
	fmt.Println("Vegy List is",vegList)
}