package main

import (
	"fmt"
)

func getit(a ...interface{}){
	fmt.Scan(a ...)
}

func printer(a ...interface{}){
	fmt.Println(a ...)
}

func make_square(a int){
	printer(a*a)
}

func pass_to_b(a int){
	make_square(a)
}

func pass_to_a(a int){
	pass_to_b(a)
}

func get_data(a int){
	pass_to_a(a)
}

func main(){
	var num int
	printer("enter a number to square: ")
	getit(&num)
	get_data(num)
}