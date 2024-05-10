package main

import "fmt"

func print(a ...interface{}){
	fmt.Println(a ...)
}

func changevalue(points *int){
	*points=56
}

func tester(add *int,num int){
	*add=num

}

func exchange(n1 *int, n2 *int){
	*nx:=n1
	*ny:=n2
	print(*nx,*ny)
	// *n1=ny
	// *n2=nx 
}

func main(){
	num:=34
	print("initial ",num)
	print("location ", &num)

	changevalue(&num)
	print("value of num after calling changevalue function: ",num)
	print("location ", &num)

	a:=100
	b:=200

	print("a,b: ",a,b)
	tester(&a,b)
	print("a,b: ",a,b)

	print("----------------------------")
	x,y:=45,54
	print("x,y: ",x,y)
	exchange(&x,&y)
	print("x,y: ",x,y)

}