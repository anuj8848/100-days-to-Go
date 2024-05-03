package main

import ("fmt"
	)

func main(){
	var i int
	fmt.Print("Type a number between 1 and 7: ")
	fmt.Scan(&i)

	switch i{
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Print("monday")
	case 3:
		fmt.Println("tuesday")
	case 4:
		fmt.Println("webnesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	}
}
