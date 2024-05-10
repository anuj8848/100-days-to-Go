package main

import "fmt"

func printString(s string){
	fmt.Println(s)
}

func main(){
	go printString("goroutines")

	fmt.Println("main function")

}