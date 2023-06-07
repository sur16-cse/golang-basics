package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	Welcome:="Welcome to userInput"
	fmt.Println(Welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")

	//comma ok //err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("Type of this rating is %T",input)
}