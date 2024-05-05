package main
import ("fmt")




func main(){
	var num int
	fmt.Print("Type a number: ")
	fmt.Scan(&num)
	if num%5==0{
		fmt.Println(num," is divisible by 5")
	}else {
		fmt.Println(num," is not divisible by 5")
	}
}

