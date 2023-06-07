package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 9 {
			break
		} else if rogueValue == 8 {
			goto lco
		} else {
			fmt.Println("Value is: ", rogueValue)
			rogueValue++
				continue

		}

	}

lco:
	fmt.Println("jumping at end")
}
