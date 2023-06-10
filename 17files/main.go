package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to files in golang")
	content:="This need to go in a file"

	file,err:=os.Create("./mygofile.text")

	if err!=nil{
		panic(err)
	}
	length,err:=io.WriteString(file,content)

	checkNilError(err)

	fmt.Println("length is:",length)
	defer file.Close()
	readFile("./mygofile.text")
}

func readFile(filname string)  {
	databyte,err:=ioutil.ReadFile(filname)
	checkNilError(err)

	fmt.Println("Text data inside the file is \n",databyte)
	fmt.Println("Text data inside the file is \n",string(databyte))
}

func checkNilError(err error)  {
	if err!=nil{
		panic(err)
	}
}