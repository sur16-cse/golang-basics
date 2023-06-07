package main

import "fmt"

const LoginToken string="bkxnxnxnkm"

func main(){
	fmt.Println("variables");
	var username string="surbhi kumari";
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username);

	var  isLoggedin bool=true;
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n",isLoggedin);

	var  smallVal uint8=255;
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal);

	var  smallFloat uint8=255;
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat);

	//default value and some aliases
	var  anotherVariable int;
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable);

	//implicit type
	var website="surbhikumaridav@gmail.com";
	fmt.Println(website);

	numberOfUser:=300000.00
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n",LoginToken);
}