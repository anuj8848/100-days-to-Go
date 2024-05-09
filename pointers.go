package main

import "fmt"

func print(a ...interface{}){
	fmt.Println(a ...)
}

func changevalue(points *int){
	*points=56
}

func main(){
	num:=34
	print("initial ",num)
	print("location ", &num)

	changevalue(&num)
	print("value of num after calling changevalue function: ",num)
	print("location ", &num)


}