package main 
import "fmt"

func print(a ...interface{}){
	fmt.Println(a ...)

}

func primechecker(num int){
	counter:=0
	for x:=2; x<=num/2;x++{
		if num%x==0{
			counter++
		}
		
	}
	if counter >0{
		print("this is not prime number")
	}
	if counter == 0{
		print("this is prime number")
	}

}

func main(){
	var num int
	print("enter a number")
	fmt.Scan(&num)

	primechecker(num)

}