package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to golang slices")
	var fruitList=[]string{"Apple","Mango","Tomato"}
	fmt.Printf("Type of fruitlist is %T\n",fruitList)

	fruitList=append(fruitList, "Peach","Banana")
	fmt.Println(fruitList)

	fruitList=append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList=append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList=append(fruitList[:1])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0]=235
	highScore[1]=232
	highScore[2]=236
	highScore[3]=231
	fmt.Println(highScore)

	highScore=append(highScore, 555,666,377)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a slice based on index value
	var courses=[]string{"reactjs","ruby","python","c","c++","dart","javascript","c#","java"}
	fmt.Println(courses)
	var index int=2
	courses=append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}